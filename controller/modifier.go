package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi-basic/helper"
	"restapi-basic/model"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Modifier interface{}

func InsertModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		log.Fatal(err)
	}
	var Modifier model.TkModifierParent
	err_cat := json.NewDecoder(r.Body).Decode(&Modifier)

	if err_cat != nil {
		log.Fatal(err_cat)
	}

	Modifier.ModifierParentId = uuid.New()

	for _, child := range Modifier.ModifierChilds {
		child.ModifierChildId = uuid.New()
	}

	validateModifiers := helper.Validate(Modifier)
	if validateModifiers != nil {
		response := helper.FailedValidate("Invalid Data", validateModifiers)
		json, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		err_insert := Modifier.InsertModifierParent(db)
		if err_insert != nil {
			log.Fatal(err_insert)
		}

		var modifierChildParents []model.TkParentChildModifiers
		for _, child := range Modifier.ModifierChilds {
			var mod_child_parent model.TkParentChildModifiers
			mod_child_parent.ModifierParentId = Modifier.ModifierParentId
			mod_child_parent.ModifierChildId = child.ModifierChildId
			modifierChildParents = append(modifierChildParents, mod_child_parent)
		}
		if modifierChildParents != nil {
			err := model.InsertModifierChildParent(db, modifierChildParents)
			if err != nil {
				log.Fatal(err)
			}
		}

		response := helper.Success(Modifier, nil, "Insert Modifier successfully")
		json, err_json := json.Marshal(response)
		if err_json != nil {
			log.Fatal(err_json)
		}
		w.Write(json)
	}

	defer helper.CloseConnection(db)
}

func SelectAllModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		log.Fatal(err)
	}
	var Modifier model.TkModifierParent

	result, err_insert := Modifier.SelectAllModifier(db)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all Modifier successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func SelectOneModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Modifier_id := params["id"]
	if err != nil {
		log.Fatal(err)
	}
	var Modifier model.TkModifierParent
	res_uuid, err := uuid.Parse(Modifier_id)
	if err != nil {
		log.Fatal(err)
	}
	Modifier.ModifierParentId = res_uuid
	result, err_insert := Modifier.SelectOneModifier(db)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	response := helper.Success(result, nil, "Select all Modifier successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}

func UpdateModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()

	if err != nil {
		log.Fatal(err)
	}
	var Modifier model.TkModifierParent
	err_cat := json.NewDecoder(r.Body).Decode(&Modifier)

	if err_cat != nil {
		log.Fatal(err_cat)
	}

	validateModifiers := helper.Validate(Modifier)
	if validateModifiers != nil {
		response := helper.FailedValidate("Invalid Data", validateModifiers)
		json, err := json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(json)
	} else {
		cat, err_update := Modifier.UpdateModifier(db)
		if err_update != nil {
			log.Fatal(err_update)
		}

		var modifierChilds []model.TkModifierChild
		var modifierChildParents []model.TkParentChildModifiers
		for _, child := range Modifier.ModifierChilds {
			var modifierChild model.TkModifierChild
			modifierChild.ModifierChildId = child.ModifierChildId
			modifierChild.ModifierChildName = child.ModifierChildName
			modifierChild.ModifierChildPrice = child.ModifierChildPrice
			modifierChild.ModifierChildDesc = child.ModifierChildDesc
			_, err := modifierChild.SelectOneModifierChild(db)
			if err != nil {
				var modifierChildParent model.TkParentChildModifiers
				modifierChildParent.ModifierParentId = Modifier.ModifierParentId
				modifierChildParent.ModifierChildId = child.ModifierChildId
				modifierChilds = append(modifierChilds, modifierChild)
				modifierChildParents = append(modifierChildParents, modifierChildParent)
			} else {
				err := modifierChild.UpdateModifierChild(db)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		err_mod_child := model.InsertModifierChild(db, modifierChilds)
		if err_mod_child != nil {
			log.Fatal(err_mod_child)
		}

		err_mod_child_parent := model.InsertModifierChildParent(db, modifierChildParents)
		if err_mod_child_parent != nil {
			log.Fatal(err_mod_child_parent)
		}

		response := helper.Success(cat, nil, "Update Modifier successfully")
		json, err_json := json.Marshal(response)
		if err_json != nil {
			log.Fatal(err_json)
		}
		w.Write(json)
	}
	defer helper.CloseConnection(db)
}

func DeleteModifier(w http.ResponseWriter, r *http.Request) {
	db, err := helper.Connection()
	params := mux.Vars(r)
	Modifier_id := params["id"]
	if err != nil {
		log.Fatal(err)
	}
	var Modifier model.TkModifierParent
	res_uuid, err := uuid.Parse(Modifier_id)
	if err != nil {
		log.Fatal(err)
	}
	Modifier.ModifierParentId = res_uuid
	err_insert := Modifier.DeleteModifier(db)
	if err_insert != nil {
		log.Fatal(err_insert)
	}

	response := helper.Success(nil, nil, "Delete Modifier successfully")
	json, err_json := json.Marshal(response)
	if err_json != nil {
		log.Fatal(err_json)
	}
	w.Write(json)
	defer helper.CloseConnection(db)
}
