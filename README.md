# API pour une clinique vétérinaire

***L'adresse pincipale locale est localhost:8080/api/v1***

## 1. Les adresses de l'API


### A. Les chats

la gestion des chats se fait sur différentes adresses :

 - **/cats** (route principale) permet soit d'afficher l'intégralité des chats enregistrés, soit d'en enregistrer un nouveau
 - **/cats/(id)** permet d'afficher un chat en particulier, modifier ses informations ou le supprimer des registres
 - **/cats/(id)/history** permet d'afficher les différentes consultations et traitements d'un chat en particulier


### B. Les visites

Pour les visites, trois possibilités s'offrent à vous :

 - **/visits** permet d'ajouter une nouvelle visite (attention bien préciser le numéro du chat conerné par le rendez-vous)
 - **/visits/(id)** permet d'afficher l'intégralité des visites pour un chat donné, ses traitements donnés lors du rendez-vous sont inclus dessus
 - **/visits/(id)/(filtre)** permet d'afficher les visites d'un chat trié dans l'ordre alphabétique selon trois critères : la date, le vétérinaire vu ou la raison du rendez-vous


### C. Les traitements

Pour les traitements, deux choses sont possibles :

 - **/treatments** permet d'ajouter un nouveau traitement (attention bien préciser le numéro du rendez-vous)
 - **/treatments/(id)** permet d'afficher tous les traitements distribués pour une visite donnée


## 2. Le fonctionnement de l'API

Pour lancer l'API, il suffit d'exécuter le programme avec 
```bash
go run main.go
```

