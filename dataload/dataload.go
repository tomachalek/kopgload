// Copyright 2018 Tomas Machalek <tomas.machalek@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dataload

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tomachalek/vertigo/v2"
)

type Loader struct {
	tx         *sql.Tx
	insertStmt *sql.Stmt
}

func NewLoader(tx *sql.Tx) *Loader {
	return &Loader{
		tx: tx,
	}
}

func (d *Loader) Prepare() {
	var err error
	d.insertStmt, err = d.tx.Prepare("INSERT INTO token (pa_word, pa_lemma, pa_tag) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func (d *Loader) Finish() {
	d.insertStmt.Close()
}

func (d *Loader) ProcToken(token *vertigo.Token, err error) {
	if err != nil {
		log.Print("ERROR: ", err)
		return
	}
	_, err = d.insertStmt.Exec(token.Word, token.Attrs[0], token.Attrs[1])
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func (d *Loader) ProcStruct(strc *vertigo.Structure, err error) {
	//strc.Attrs
}

func (d *Loader) ProcStructClose(strc *vertigo.StructureClose, err error) {

}
