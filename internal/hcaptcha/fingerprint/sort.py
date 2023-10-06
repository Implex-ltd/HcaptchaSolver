import json

# Les données initiales
data = {
   "fingerprint_events": [
        [
            3,
            "8683.100000023842"
        ],
        [
            1902,
            "57"
        ],
        [
            1901,
            "15307345790125003576"
        ],
        # ... (toutes les autres données)
    ]
}

# Tri des données en fonction de la première valeur de chaque sous-liste
sorted_data = sorted(data["fingerprint_events"], key=lambda x: x[0])

# Création d'un nouveau dictionnaire trié
sorted_data_dict = {"fingerprint_events": sorted_data}

# Conversion en JSON trié
sorted_json = json.dumps(sorted_data_dict, indent=4)

# Imprimer le JSON trié
print(sorted_json)
