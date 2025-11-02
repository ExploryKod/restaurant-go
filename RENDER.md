# Configuration Render.com

Ce document explique comment configurer l'application RestaurantGo sur Render.com avec la base de données Aiven MySQL.

## Variables d'environnement à configurer dans Render.com

**IMPORTANT** : Les variables d'environnement doivent être définies dans la section **"Environment"** (pas seulement dans "Secrets") pour que Render.com les injecte dans le conteneur Docker.

### Base de données (Aiven MySQL)

Dans le dashboard Render.com de votre service web, allez dans **Settings > Environment** et ajoutez :

```
BDD_USER=avnadmin
BDD_PASSWORD=<votre_mot_de_passe_aiven>
BDD_PORT=mysql-restaurantgo-amauryfranssen-5ab7.h.aivencloud.com:14253
BDD_NAME=defaultdb
```

**Important** : 
- `BDD_PORT` doit contenir **host:port** ensemble (pas juste le port)
- Le mot de passe Aiven doit être récupéré depuis votre dashboard Aiven
- Pour `BDD_PASSWORD`, vous pouvez utiliser la fonctionnalité "Secrets" de Render.com, mais assurez-vous qu'elle est aussi référencée dans "Environment Variables"

### Port de l'application

L'application écoute sur le port **9999** par défaut. Render.com devrait automatiquement détecter ce port depuis le Dockerfile (`EXPOSE 9999`).

## Configuration Render.com

### Option 1 : Utiliser le fichier `render.yaml` (Recommandé)

Le fichier `render.yaml` à la racine du projet définit la configuration du service. Render.com le détectera automatiquement lors du déploiement.

**Note** : `BDD_PASSWORD` doit être défini manuellement dans les "Secrets" de Render.com pour des raisons de sécurité (il n'est pas dans le fichier `render.yaml`).

### Option 2 : Configuration manuelle dans le dashboard

1. **Service Type** : Web Service
2. **Environment** : Docker
3. **Dockerfile Path** : `./Dockerfile`
4. **Docker Context** : `.` (racine du projet)
5. **Build Command** : Laissé vide (Dockerfile gère tout)
6. **Start Command** : Laissé vide (ENTRYPOINT dans Dockerfile)
7. **Environment Variables** : Ajoutez manuellement toutes les variables BDD_*

## Étapes de déploiement

1. **Configurer les variables d'environnement** :
   - Allez dans votre service web sur Render.com
   - Ouvrez **Settings > Environment**
   - Ajoutez toutes les variables `BDD_*` dans "Environment Variables"
   - Pour `BDD_PASSWORD`, utilisez la fonctionnalité "Secrets" pour plus de sécurité

2. **Vérifier le Dockerfile** :
   - Assurez-vous que le `Dockerfile` (production) est utilisé, pas `dev.Dockerfile`
   - Le `.dockerignore` exclut le fichier `.env` du build

3. **Déployer** :
   - Poussez vos changements sur Git
   - Render.com détectera automatiquement le fichier `render.yaml` (si présent)
   - Ou lancez un déploiement manuel depuis le dashboard

## Vérification

Après le déploiement, vérifiez les logs pour confirmer :
1. La détection de l'environnement Render.com
2. La connexion à la base de données

Vous devriez voir dans les logs :
```
=== Database Configuration Debug ===
Is Render.com environment: true
BDD_USER: avnadmin
BDD_PORT: mysql-restaurantgo-amauryfranssen-5ab7.h.aivencloud.com:14253
BDD_NAME: defaultdb
BDD_PASSWORD: ***SET***
Connecting to database: user=avnadmin, addr=mysql-restaurantgo-amauryfranssen-5ab7.h.aivencloud.com:14253, dbname=defaultdb
```

Si vous voyez `user=, addr=, dbname=` dans les logs, cela signifie que les variables d'environnement ne sont pas correctement injectées. Vérifiez qu'elles sont bien dans "Environment Variables" et non seulement dans "Secrets".

## Séparation Dev/Prod

- **Développement local** : Utilise `dev.Dockerfile` et `docker-compose.yaml`
- **Production Render.com** : Utilise `Dockerfile` (défini dans `render.yaml` ou dans les settings du service)
- Le code détecte automatiquement l'environnement Render.com et ne charge jamais le `.env` en production

