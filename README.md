Go STAR Bus 83 Live

Application Go permettant de consulter en temps réel les prochains passages de la ligne de bus 83 du réseau STAR (Rennes), en utilisant les API Open Data de Rennes Métropole.

Description

Ce projet interroge l'API temps réel du réseau STAR afin de récupérer les horaires de passage du bus 83 aux différents arrêts du réseau.

Les données sont récupérées au format JSON, puis converties en structures Go pour être traitées et affichées dans le terminal.

Fonctionnalités
Connexion à l'API Open Data STAR
Récupération des données en temps réel
Désérialisation du JSON avec encoding/json
Filtrage des passages de la ligne 83
Affichage des arrêts desservis
Affichage des destinations
Affichage des heures d'arrivée prévues
Comptage du nombre de passages trouvés
Gestion des erreurs HTTP et JSON
Utilisation d'un client HTTPS personnalisé pour ignorer les erreurs TLS dans un environnement de test
