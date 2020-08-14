/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package jsonast

import (
	"net/url"
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ObjectTypeOf(t *testing.T) {
	g := NewGomegaWithT(t)

	url, err := url.Parse("https://schema.management.azure.com/schemas/2015-01-01/Microsoft.Resources.json#/resourceDefinitions/deployments")
	g.Expect(err).To(BeNil())

	name := objectTypeOf(url)
	g.Expect(name).To(Equal("deployments"))
}
