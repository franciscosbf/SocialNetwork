# DB Descrition

### Accounts

- Main and personal info of user.
- *username* and *email* must be unique.
- each pair (*username*, *email*) must be unique.
- Contains attributes related to security and privacy.

### Profiles

- Public info which describes the user.
- Address is composed by latitude and longitude and created depending on some external service which provides geodeconding based in a given address.

### Hierarchy between *Accounts* and *Profiles*

- *Profiles* extends *Accounts* to keep personal and public info separated. However, a user profile has the same uid as the account, since its inherited.

### Description of Entity Relationship 

- accounts (**aid**, username, password, email, phone_prefix, phone_number)
- profiles (**pid**, aid, first_name, second_name, surname, description, address, addr_latitude, addr_longitude)
    - aid is foreign key, referring accounts
