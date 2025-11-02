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

### Configuration pour la production (Render.com)

Pour le déploiement sur Render.com avec Aiven MySQL, configurez ces variables d'environnement dans les **Secrets** de Render.com :

```
BDD_USER=avnadmin
BDD_PASSWORD=<votre_mot_de_passe_aiven>
BDD_PORT=mysql-restaurantgo-amauryfranssen-5ab7.h.aivencloud.com:14253
BDD_NAME=defaultdb
```

**Important** : `BDD_PORT` doit contenir **host:port** ensemble (format: `hostname:port`)

Voir `RENDER.md` pour plus de détails sur le déploiement sur Render.com.

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
