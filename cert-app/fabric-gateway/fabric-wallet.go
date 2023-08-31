package main

// import (
// 	"encoding/json"

// 	"github.com/hyperledger/fabric-gateway/pkg/identity"
// 	"github.com/pkg/errors"
// )

// // A Wallet stores identity information used to connect to a Hyperledger Fabric network.
// // Instances are created using factory methods on the implementing objects.
// type Wallet struct {
// }

// // Put an identity into the wallet
// //
// //	Parameters:
// //	label specifies the name to be associated with the identity.
// //	id specifies the identity to store in the wallet.
// func (w *Wallet) Put(label string, id identity.X509Identity) error {
// 	content := id.Credentials() // Returns an X.509 certificate in byte

// 	return
// }

// // Get an identity from the wallet. The implementation class of the identity object will vary depending on its type.
// //
// //	Parameters:
// //	label specifies the name of the identity in the wallet.
// //
// //	Returns:
// //	The identity object.
// func (w *Wallet) Get(label string) (Identity, error) {
// 	content, err := w.store.Get(label)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var data map[string]interface{}
// 	if err := json.Unmarshal(content, &data); err != nil {
// 		return nil, errors.Wrap(err, "Invalid identity format")
// 	}

// 	idType, ok := data["type"].(string)

// 	if !ok {
// 		return nil, errors.New("Invalid identity format: missing type property")
// 	}

// 	var id Identity

// 	switch idType {
// 	case x509Type:
// 		id = &X509Identity{}
// 	default:
// 		return nil, errors.New("Invalid identity format: unsupported identity type: " + idType)
// 	}

// 	return id.fromJSON(content)
// }

// // List returns the labels of all identities in the wallet.
// //
// //	Returns:
// //	A list of identity labels in the wallet.
// func (w *Wallet) List() ([]string, error) {
// 	return w.store.List()
// }

// // Exists tests whether the wallet contains an identity for the given label.
// //
// //	Parameters:
// //	label specifies the name of the identity in the wallet.
// //
// //	Returns:
// //	True if the named identity is in the wallet.
// func (w *Wallet) Exists(label string) bool {
// 	return w.store.Exists(label)
// }

// // Remove an identity from the wallet. If the identity does not exist, this method does nothing.
// //
// //	Parameters:
// //	label specifies the name of the identity in the wallet.
// func (w *Wallet) Remove(label string) error {
// 	return w.store.Remove(label)
// }
