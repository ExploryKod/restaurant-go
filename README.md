# restaurant-go

Nous n'avons pas pu merger tout sur main ni dev donc pour aller voir les fonctionnalités, ne pas hesiter à regarder les branches.

## Installation 

Demandez-nous le reste du .env file (transmis lors de la soutenance).

Le .env pour une partie des variables: 
```
GOOS=linux
GOARCH=amd64
SERVER_URL=localhost:8097
BDD_PASSWORD=password
BDD_USER=root
BDD_NAME=restaurantbdd
BDD_PORT=127.0.0.1:3309
```

### Email:
Ouvrez un compte gratuit MailTrap de test et aller dans Email Testing (colone à gauche) > SMTP settings > cliquer sur "show credentials"

Coller vos credentials dans le .env : 
```
SMTP_PASSWORD=<votrepassword>
SMTP_USERNAME=<usernamemailtrap>
SMTP_HOST=sandbox.smtp.mailtrap.io
SMTP_PORT=25
```
1. Base de donnée et phpmyadmin 

``` docker compose up -d --build ```

2. Démarrer l'application

``` go mod tidy ```

``` go run main/main.go ```

Aller sur ```http://localhost:9999/login```

2. Alternative : utiliser air pour le hot reloading

Installer et configurer air : [Visitez leur site](https://github.com/air-verse/air)

**Installer:**
```go
go install github.com/air-verse/air@latest
```
**Configurer les fichier de configuration :** <br>
(la commande lancera air si elle trouve un fichier déjà présent)
```go
air -c .air.toml
```

**Executer avec simplement:**
``` air ```
