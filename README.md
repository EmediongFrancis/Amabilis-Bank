# Amabilis Bank ![Bank Icon](https://github.com/EmediongFrancis/Amabilis-Bank/blob/main/assets/Bank.png)

Amabilis Bank is the backend service of a basic banking system built with PostgreSQL, Golang, and Docker. This bank will provide APIs for the frontend to:
- Create and manage accounts:
  - Accounts will be composed of owner&rsquo;s name, account balance, and banking currency.
- Record all balance changes:
  - Account entries will be created for every change in balance.
- Perform money transfer:
  - Transfer will be done in a consistent manner within a transaction such that both account balances (sender and receiver) are updated simultaneously if the transfer is successful or none of them are updated in the event of a failed transfer. <br><br>

| ![Database Schema](https://github.com/EmediongFrancis/Amabilis-Bank/blob/main/assets/Amabilis%20Bank.png) |
|:--:|
| <b> Database Schema </b>

## Tools & Technologies
* <b>Golang</b> - Primary Language.
* <b>PostgreSQL</b> - Database.
* <b>Docker</b> - Containerization.
* <b>Git</b> - Version Control.
* <b>Makefile</b> - Build System Efficiency.
* <b>TablePlus</b> - Database Management.
* <b>SQLc</b> - Golang's SQL Compiler Tool (Compiles PostgreSQL code into native Golang).
* <b>Golang-Testify</b> - Testing Framework.

<h3><i>Work in progress...</i></h3>
