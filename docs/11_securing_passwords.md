# Securing Passwords 

* Steps for securing passords 
* Third party auth options 
* Hashing data 
* Storing password hashes 
* Using password salt 

## Steps for securing passords 

Current identity: an id, an email address, a pwd. 
Practices to follow: 
1. Use HTTPS 
2. Store hashed pwds. Never encrypted or plaintext ones.
3. Add salt to pwds
4. Used time constant functions during auth. 

Third party auth options: 
* Libraries like `devise` for RoR. 
* Saas services like Auth0. 

## Using password salt 

Example: 
* User pwd: `abc123`
* We add a random salt: `ja08d`
* We hash the following: `abc123-ja08d`
* Resulting hash: `347563...` (to be stored in db)
