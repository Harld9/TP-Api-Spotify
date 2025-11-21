🎧 TP API Spotify

📌 Objectif

Créer un petit site web en Go qui affiche des données récupérées depuis l’API Spotify en utilisant :
un serveur HTTP
des templates GoHTML
des requêtes API directement dans les handlers
une gestion automatique du token Spotify

🚀 Fonctionnalités

🔹 Route /album/damso

Afficher la liste des albums de Damso avec :
Nom de l’album
Image de couverture
Date de sortie
Nombre de musiques

🔹 Route /track/laylow

Afficher les informations du titre “Maladresse” de Laylow :

Nom du titre
Image de l’album
Nom de l’album
Artiste
Date de sortie
Lien Spotify direct

⚙️ Contraintes techniques

Le serveur écoute sur le port 8080
Les appels à l’API sont effectués dans les handlers
Le token doit être régénéré avant expiration
Le design du site est libre mais doit être propre (mise en page, couleurs, polices)

🔗 Dépôt GitHub

👉 https://github.com/Harld9/TP-Api-Spotify
