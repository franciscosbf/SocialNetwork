# DB Descrition

### Accounts

- Main and personal info of user.
- Contains attributes related to security and privacy.

### Profiles

- Public info which describes the user.
- Address is composed by latitude and longitude and created depending on some external service which provides geodeconding based in a given address.

### Hierarchy between *Accounts* and *Profiles*

- *Profiles* extends *Accounts* to keep personal and public info separated. However, a user profile has the same username as the account, since its inherited.

### Description of Entity Relationship 

- accounts (**username**, password, email)
- profiles (**username**, first_name, second_name, surname, description, address, latitude, longitude)
    - username is the same pk as accounts
