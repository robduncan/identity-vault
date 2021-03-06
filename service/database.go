// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016-2017 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package service

import (
	"database/sql"
	"log"

	// postgresql driver
	_ "github.com/lib/pq"
)

// Datastore interface for the database logic
type Datastore interface {
	ListModels() ([]Model, error)
	FindModel(brandID, modelName string) (Model, error)
	GetModel(modelID int) (Model, error)
	UpdateModel(model Model) (string, error)
	DeleteModel(model Model) (string, error)
	CreateModel(model Model) (Model, string, error)
	CreateModelTable() error

	ListKeypairs() ([]Keypair, error)
	GetKeypair(keypairID int) (Keypair, error)
	PutKeypair(keypair Keypair) (string, error)
	UpdateKeypairActive(keypairID int, active bool) error
	CreateKeypairTable() error

	CreateSettingsTable() error
	PutSetting(setting Setting) error
	GetSetting(code string) (Setting, error)

	CreateSigningLogTable() error
	CheckForDuplicate(signLog *SigningLog) (bool, int, error)
	CreateSigningLog(signLog SigningLog) error
	ListSigningLog(fromID int) ([]SigningLog, error)
	DeleteSigningLog(signingLog SigningLog) (string, error)

	CreateDeviceNonceTable() error
	CreateDeviceNonce() (DeviceNonce, error)
	ValidateDeviceNonce(nonce string) error
}

// DB local database interface with our custom methods.
type DB struct {
	*sql.DB
}

// OpenSysDatabase Return an open database connection
func OpenSysDatabase(driver, dataSource string) *DB {
	// Open the database connection
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatalf("Error opening the database: %v\n", err)
	}

	// Check that we have a valid database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error accessing the database: %v\n", err)
	} else {
		log.Println("Database opened successfully.")
	}
	return &DB{db}
}
