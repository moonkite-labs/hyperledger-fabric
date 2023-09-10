#!/bin/bash
export PATH=${PWD}/bin:$PATH
export FABRIC_CFG_PATH=$PWD/config/

function asOrg() {
    echo "Using ${1}"
    ORDERER_CA=${PWD}/test-network/organizations/ordererOrganizations/example.com/tlsca/tlsca.example.com-cert.pem
    PEER0_ORG_CA=${PWD}/test-network/organizations/peerOrganizations/${1}.example.com/tlsca/tlsca.${1}.example.com-cert.pem
    CORE_PEER_LOCALMSPID=${1^}MSP
    CORE_PEER_MSPCONFIGPATH=${PWD}/test-network/organizations/peerOrganizations/${1}.example.com/users/User1@${1}.example.com/msp
    case "${1}" in
        org1 )
            CORE_PEER_ADDRESS=localhost:7051
            ;;
        org2 )
            CORE_PEER_ADDRESS=localhost:9051
            ;;
        org3 )
            CORE_PEER_ADDRESS=localhost:11051
            ;;
    esac
    CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/test-network/organizations/peerOrganizations/${1}.example.com/tlsca/tlsca.${1}.example.com-cert.pem

    export CORE_PEER_TLS_ENABLED=true
    export ORDERER_CA="${ORDERER_CA}"
    export PEER0_ORG_CA="${PEER0_ORG_CA}"

    export CORE_PEER_MSPCONFIGPATH="${CORE_PEER_MSPCONFIGPATH}"
    export CORE_PEER_ADDRESS="${CORE_PEER_ADDRESS}"
    export CORE_PEER_TLS_ROOTCERT_FILE="${CORE_PEER_TLS_ROOTCERT_FILE}"

    export CORE_PEER_LOCALMSPID="${CORE_PEER_LOCALMSPID}"
}

function parsePeerConnectionParameters() {
    PEER_CONN_PARMS=()
    PEERS=""
    while [ "$#" -gt 0 ]; do
        asOrg "org$1"
        PEER="peer0.org$1"
        ## Set peer addresses
        if [ -z "$PEERS" ]
        then
            PEERS="$PEER"
        else
            PEERS="$PEERS $PEER"
        fi
        PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
        ## Set path to TLS certificate
        CA=PEER0_ORG_CA
        TLSINFO=(--tlsRootCertFiles "${!CA}")
        PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
        # shift by one to get to the next organization
        shift
    done
}

function parseArgs() {

    if [[ -z "$3" ]]; then
        ARGS=''
    else
        ARGS=", $(echo $3 | sed 's/[^[:space:],]\+/"&"/g')"
    fi

    if [[ -z "$1" ]]; then
        FUNC_NAME="$2"
    else
        FUNC_NAME="$1:$2"
    fi
}

function invoke() {
    parseArgs "$@"

    local OLD_ORG=$ORG
    parsePeerConnectionParameters 1 2 3
    asOrg $OLD_ORG
    set -x;
    peer chaincode invoke --orderer "0.0.0.0:7050" --ordererTLSHostnameOverride orderer.example.com --tls --cafile "$ORDERER_CA" -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" "${PEER_CONN_PARMS[@]}" -c "{\"Args\":[\"${FUNC_NAME}\"${ARGS}]}"
    { set +x; } 2>/dev/null
}

function query() {
    asOrg "$ORG"
    parseArgs "$@"

    echo "CONTRACT_NAME: $1, FUNCTION_NAME: $2, ARGS: $3"

    set -x;
    peer chaincode query --orderer "0.0.0.0:7050" --ordererTLSHostnameOverride orderer.example.com --tls --cafile "$ORDERER_CA" -C "$CHANNEL_NAME" -n "$CHAINCODE_NAME" -c "{\"Args\":[\"${FUNC_NAME}\"${ARGS}]}"
    { set +x; } 2>/dev/null
}

function printHelp() {
    println "Usage: `basename ${0}` <command> [args]"
    println "  Commands:"
    println "    invoke - To invoke the chaincode as a peer."
    println "    query - To query the chaincode as a peer." 
    println "  Flags:" 
    println "    -n <channel_name>      Channel name to connect the interact with the chaincode"
    println "    -g <chaincode_name>    Chaincode name to be used. Default is 'checkpoints'"
    println "    -c <contract_name>     Contract name to be used."
    println "    -f <func_name>         Function name to be called." 
    println "    -u <org_name>          Organization name (e.g. org1, org2...)." 
    println "    -a <args>              Arguments to be passed into a function, ignore if None." 
    println "    -h                     Print this help message."
    println "  Example: ./`basename ${0}` query -n certapp -g certificate-manager -c CertificateContract -f GetAllCertificates -u org1"
    println "           ./`basename ${0}` invoke -n certapp -g certificate-manager -c CertificateContract -f IssueCertificate -u org1 -a \"arguments separated with space...\""
}

function println() {
    echo -e "$1"
}

if [[ $# -lt 1 ]]; then
    printHelp
    exit 1
else    
    COMMAND=$1
    shift
fi

while getopts ":n:g:c:f:u:a:" opt; do
    case ${opt} in
    n ) 
        CHANNEL_NAME=$OPTARG
        ;;
    g ) 
        CHAINCODE_NAME=$OPTARG
        ;;
    c )
        CONTRACT_NAME=$OPTARG
        ;;
    f )
        FUNCTION_NAME=$OPTARG
        ;;
    u ) 
        ORG=${OPTARG:-'org1'}
        ;;
    a ) 
        FUNC_ARGS+=("$OPTARG")
        ;;
    h ) 
        printHelp
        exit 1
        ;;
    \? )
        echo "Invalid Option: -$OPTARG" 1>&2
        exit 1
        ;;
    : )
        echo "Invalid Option: -$OPTARG requires an argument" 1>&2
        exit 1
        ;;
    esac
done
shift $((OPTIND -1))

# : ${CONTRACT_NAME:?'Missing contract name -c!'} ${FUNCTION_NAME:?'Missing function name -f!'}

case "$COMMAND" in
    invoke)
        invoke "$CONTRACT_NAME" "$FUNCTION_NAME" "${FUNC_ARGS[@]}"
        ;;
    query)
        query "$CONTRACT_NAME" "$FUNCTION_NAME" "${FUNC_ARGS[@]}"
        ;;
    * )
        printHelp
        exit 1
        ;;
esac