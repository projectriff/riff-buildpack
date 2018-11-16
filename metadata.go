/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package riff_buildpack

import (
	"fmt"
	"path/filepath"

	"github.com/buildpack/libbuildpack"
	"github.com/cloudfoundry/libjavabuildpack"
)

// Metadata represents the contents of the riff.toml file in an application root
type Metadata struct {
	// Artifact is the path to the main function artifact. This may be a java jar file, an executable file, etc
	// May be autodetected or chosen by a collaborating buildpack
	Artifact string `toml:"artifact"`

	// Handler is a "finer grained" handler for the function within the artifact, if applicable.
	// This may be a classname, a function name, etc. May be autodetected or chosen by a collaborating
	// buildpack or function invoker.
	Handler string `toml:"handler"`

	// Override is an optional value provided by the user to force a given language for the function and
	// completely bypass the detection mechanism, if needed.
	Override string `toml:"override"`
}

// String makes Metadata satisfy the Stringer interface.
func (m Metadata) String() string {
	return fmt.Sprintf("Metadata{ Artifact: %s, Handler: %s }", m.Artifact, m.Handler)
}

// NewMetadata creates a new Metadata from the contents of $APPLICATION_ROOT/riff.toml.
func NewMetadata(application libbuildpack.Application, logger libjavabuildpack.Logger) (Metadata, error) {
	f := filepath.Join(application.Root, "riff.toml")

	exists, err := libjavabuildpack.FileExists(f)
	if err != nil {
		return Metadata{}, err
	}

	if !exists {
		return Metadata{}, nil
	}

	var metadata Metadata
	err = libjavabuildpack.FromTomlFile(f, &metadata)
	if err != nil {
		return Metadata{}, err
	}

	logger.Debug("riff metadata: %s", metadata)
	return metadata, nil
}
