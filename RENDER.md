# Configuration Render.com

Ce document explique comment configurer l'application RestaurantGo sur Render.com avec la base de données Aiven MySQL.

## Variables d'environnement à configurer dans Render.com

Dans les **Secrets** de votre service Render.com, configurez les variables suivantes :

### Base de données (Aiven MySQL)

```
BDD_USER=avnadmin
BDD_PASSWORD=<votre_mot_de_passe_aiven>
BDD_PORT=mysql-restaurantgo-amauryfranssen-5ab7.h.aivencloud.com:14253
BDD_NAME=defaultdb
```

**Important** : 
- `BDD_PORT` doit contenir **host:port** ensemble (pas juste le port)
- Le mot de passe Aiven doit être récupéré depuis votre dashboard Aiven

### Port de l'application

L'application écoute sur le port **9999** par défaut. Render.com devrait automatiquement détecter ce port depuis le Dockerfile (`EXPOSE 9999`).

## Configuration Render.com

1. **Service Type** : Web Service
2. **Dockerfile** : Utilisez le fichier `Dockerfile` à la racine du projet
3. **Build Command** : Laissé vide (Dockerfile gère tout)
4. **Start Command** : Laissé vide (ENTRYPOINT dans Dockerfile)

## Vérification

Après le déploiement, vérifiez les logs pour confirmer la connexion à la base de données. Vous devriez voir :
```
Connecting to database: user=avnadmin, addr=mysql-restaurantgo-amauryfranssen-5ab7.h.aivencloud.com:14253, dbname=defaultdb
```

