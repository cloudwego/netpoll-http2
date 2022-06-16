// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import "testing"

func TestIsBadCipherBad(t *testing.T) {
	for _, c := range badCiphers {
		if !isBadCipher(c) {
			t.Errorf("Wrong result for isBadCipher(%d), want true", c)
		}
	}
}

// verify we don't give false positives on ciphers not on blacklist
func TestIsBadCipherGood(t *testing.T) {
	goodCiphers := map[uint16]string{
		cipher_TLS_DHE_RSA_WITH_AES_256_CCM:                "cipher_TLS_DHE_RSA_WITH_AES_256_CCM",
		cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CCM:            "cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CCM",
		cipher_TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256: "cipher_TLS_ECDHE_PSK_WITH_CHACHA20_POLY1305_SHA256",
	}
	for c, name := range goodCiphers {
		if isBadCipher(c) {
			t.Errorf("Wrong result for isBadCipher(%d) %s, want false", c, name)
		}
	}
}

// copied from https://http2.github.io/http2-spec/#BadCipherSuites,
var badCiphers = []uint16{
	cipher_TLS_NULL_WITH_NULL_NULL,
	cipher_TLS_RSA_WITH_NULL_MD5,
	cipher_TLS_RSA_WITH_NULL_SHA,
	cipher_TLS_RSA_EXPORT_WITH_RC4_40_MD5,
	cipher_TLS_RSA_WITH_RC4_128_MD5,
	cipher_TLS_RSA_WITH_RC4_128_SHA,
	cipher_TLS_RSA_EXPORT_WITH_RC2_CBC_40_MD5,
	cipher_TLS_RSA_WITH_IDEA_CBC_SHA,
	cipher_TLS_RSA_EXPORT_WITH_DES40_CBC_SHA,
	cipher_TLS_RSA_WITH_DES_CBC_SHA,
	cipher_TLS_RSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_DH_DSS_EXPORT_WITH_DES40_CBC_SHA,
	cipher_TLS_DH_DSS_WITH_DES_CBC_SHA,
	cipher_TLS_DH_DSS_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_DH_RSA_EXPORT_WITH_DES40_CBC_SHA,
	cipher_TLS_DH_RSA_WITH_DES_CBC_SHA,
	cipher_TLS_DH_RSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_DHE_DSS_EXPORT_WITH_DES40_CBC_SHA,
	cipher_TLS_DHE_DSS_WITH_DES_CBC_SHA,
	cipher_TLS_DHE_DSS_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_DHE_RSA_EXPORT_WITH_DES40_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_DES_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_DH_anon_EXPORT_WITH_RC4_40_MD5,
	cipher_TLS_DH_anon_WITH_RC4_128_MD5,
	cipher_TLS_DH_anon_EXPORT_WITH_DES40_CBC_SHA,
	cipher_TLS_DH_anon_WITH_DES_CBC_SHA,
	cipher_TLS_DH_anon_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_KRB5_WITH_DES_CBC_SHA,
	cipher_TLS_KRB5_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_KRB5_WITH_RC4_128_SHA,
	cipher_TLS_KRB5_WITH_IDEA_CBC_SHA,
	cipher_TLS_KRB5_WITH_DES_CBC_MD5,
	cipher_TLS_KRB5_WITH_3DES_EDE_CBC_MD5,
	cipher_TLS_KRB5_WITH_RC4_128_MD5,
	cipher_TLS_KRB5_WITH_IDEA_CBC_MD5,
	cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_SHA,
	cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_SHA,
	cipher_TLS_KRB5_EXPORT_WITH_RC4_40_SHA,
	cipher_TLS_KRB5_EXPORT_WITH_DES_CBC_40_MD5,
	cipher_TLS_KRB5_EXPORT_WITH_RC2_CBC_40_MD5,
	cipher_TLS_KRB5_EXPORT_WITH_RC4_40_MD5,
	cipher_TLS_PSK_WITH_NULL_SHA,
	cipher_TLS_DHE_PSK_WITH_NULL_SHA,
	cipher_TLS_RSA_PSK_WITH_NULL_SHA,
	cipher_TLS_RSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA,
	cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_DH_anon_WITH_AES_128_CBC_SHA,
	cipher_TLS_RSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_DH_DSS_WITH_AES_256_CBC_SHA,
	cipher_TLS_DH_RSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_DHE_DSS_WITH_AES_256_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_DH_anon_WITH_AES_256_CBC_SHA,
	cipher_TLS_RSA_WITH_NULL_SHA256,
	cipher_TLS_RSA_WITH_AES_128_CBC_SHA256,
	cipher_TLS_RSA_WITH_AES_256_CBC_SHA256,
	cipher_TLS_DH_DSS_WITH_AES_128_CBC_SHA256,
	cipher_TLS_DH_RSA_WITH_AES_128_CBC_SHA256,
	cipher_TLS_DHE_DSS_WITH_AES_128_CBC_SHA256,
	cipher_TLS_RSA_WITH_CAMELLIA_128_CBC_SHA,
	cipher_TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA,
	cipher_TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA,
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA,
	cipher_TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_AES_128_CBC_SHA256,
	cipher_TLS_DH_DSS_WITH_AES_256_CBC_SHA256,
	cipher_TLS_DH_RSA_WITH_AES_256_CBC_SHA256,
	cipher_TLS_DHE_DSS_WITH_AES_256_CBC_SHA256,
	cipher_TLS_DHE_RSA_WITH_AES_256_CBC_SHA256,
	cipher_TLS_DH_anon_WITH_AES_128_CBC_SHA256,
	cipher_TLS_DH_anon_WITH_AES_256_CBC_SHA256,
	cipher_TLS_RSA_WITH_CAMELLIA_256_CBC_SHA,
	cipher_TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA,
	cipher_TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA,
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA,
	cipher_TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA,
	cipher_TLS_PSK_WITH_RC4_128_SHA,
	cipher_TLS_PSK_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_PSK_WITH_AES_128_CBC_SHA,
	cipher_TLS_PSK_WITH_AES_256_CBC_SHA,
	cipher_TLS_DHE_PSK_WITH_RC4_128_SHA,
	cipher_TLS_DHE_PSK_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_DHE_PSK_WITH_AES_128_CBC_SHA,
	cipher_TLS_DHE_PSK_WITH_AES_256_CBC_SHA,
	cipher_TLS_RSA_PSK_WITH_RC4_128_SHA,
	cipher_TLS_RSA_PSK_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_RSA_PSK_WITH_AES_128_CBC_SHA,
	cipher_TLS_RSA_PSK_WITH_AES_256_CBC_SHA,
	cipher_TLS_RSA_WITH_SEED_CBC_SHA,
	cipher_TLS_DH_DSS_WITH_SEED_CBC_SHA,
	cipher_TLS_DH_RSA_WITH_SEED_CBC_SHA,
	cipher_TLS_DHE_DSS_WITH_SEED_CBC_SHA,
	cipher_TLS_DHE_RSA_WITH_SEED_CBC_SHA,
	cipher_TLS_DH_anon_WITH_SEED_CBC_SHA,
	cipher_TLS_RSA_WITH_AES_128_GCM_SHA256,
	cipher_TLS_RSA_WITH_AES_256_GCM_SHA384,
	cipher_TLS_DH_RSA_WITH_AES_128_GCM_SHA256,
	cipher_TLS_DH_RSA_WITH_AES_256_GCM_SHA384,
	cipher_TLS_DH_DSS_WITH_AES_128_GCM_SHA256,
	cipher_TLS_DH_DSS_WITH_AES_256_GCM_SHA384,
	cipher_TLS_DH_anon_WITH_AES_128_GCM_SHA256,
	cipher_TLS_DH_anon_WITH_AES_256_GCM_SHA384,
	cipher_TLS_PSK_WITH_AES_128_GCM_SHA256,
	cipher_TLS_PSK_WITH_AES_256_GCM_SHA384,
	cipher_TLS_RSA_PSK_WITH_AES_128_GCM_SHA256,
	cipher_TLS_RSA_PSK_WITH_AES_256_GCM_SHA384,
	cipher_TLS_PSK_WITH_AES_128_CBC_SHA256,
	cipher_TLS_PSK_WITH_AES_256_CBC_SHA384,
	cipher_TLS_PSK_WITH_NULL_SHA256,
	cipher_TLS_PSK_WITH_NULL_SHA384,
	cipher_TLS_DHE_PSK_WITH_AES_128_CBC_SHA256,
	cipher_TLS_DHE_PSK_WITH_AES_256_CBC_SHA384,
	cipher_TLS_DHE_PSK_WITH_NULL_SHA256,
	cipher_TLS_DHE_PSK_WITH_NULL_SHA384,
	cipher_TLS_RSA_PSK_WITH_AES_128_CBC_SHA256,
	cipher_TLS_RSA_PSK_WITH_AES_256_CBC_SHA384,
	cipher_TLS_RSA_PSK_WITH_NULL_SHA256,
	cipher_TLS_RSA_PSK_WITH_NULL_SHA384,
	cipher_TLS_RSA_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_DH_DSS_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_DH_RSA_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_DH_anon_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_RSA_WITH_CAMELLIA_256_CBC_SHA256,
	cipher_TLS_DH_DSS_WITH_CAMELLIA_256_CBC_SHA256,
	cipher_TLS_DH_RSA_WITH_CAMELLIA_256_CBC_SHA256,
	cipher_TLS_DHE_DSS_WITH_CAMELLIA_256_CBC_SHA256,
	cipher_TLS_DHE_RSA_WITH_CAMELLIA_256_CBC_SHA256,
	cipher_TLS_DH_anon_WITH_CAMELLIA_256_CBC_SHA256,
	cipher_TLS_EMPTY_RENEGOTIATION_INFO_SCSV,
	cipher_TLS_ECDH_ECDSA_WITH_NULL_SHA,
	cipher_TLS_ECDH_ECDSA_WITH_RC4_128_SHA,
	cipher_TLS_ECDH_ECDSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_ECDHE_ECDSA_WITH_NULL_SHA,
	cipher_TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
	cipher_TLS_ECDHE_ECDSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_ECDH_RSA_WITH_NULL_SHA,
	cipher_TLS_ECDH_RSA_WITH_RC4_128_SHA,
	cipher_TLS_ECDH_RSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_ECDH_RSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_ECDH_RSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_ECDHE_RSA_WITH_NULL_SHA,
	cipher_TLS_ECDHE_RSA_WITH_RC4_128_SHA,
	cipher_TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_ECDH_anon_WITH_NULL_SHA,
	cipher_TLS_ECDH_anon_WITH_RC4_128_SHA,
	cipher_TLS_ECDH_anon_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_ECDH_anon_WITH_AES_128_CBC_SHA,
	cipher_TLS_ECDH_anon_WITH_AES_256_CBC_SHA,
	cipher_TLS_SRP_SHA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_SRP_SHA_RSA_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_SRP_SHA_DSS_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_SRP_SHA_WITH_AES_128_CBC_SHA,
	cipher_TLS_SRP_SHA_RSA_WITH_AES_128_CBC_SHA,
	cipher_TLS_SRP_SHA_DSS_WITH_AES_128_CBC_SHA,
	cipher_TLS_SRP_SHA_WITH_AES_256_CBC_SHA,
	cipher_TLS_SRP_SHA_RSA_WITH_AES_256_CBC_SHA,
	cipher_TLS_SRP_SHA_DSS_WITH_AES_256_CBC_SHA,
	cipher_TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
	cipher_TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,
	cipher_TLS_ECDH_ECDSA_WITH_AES_128_CBC_SHA256,
	cipher_TLS_ECDH_ECDSA_WITH_AES_256_CBC_SHA384,
	cipher_TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
	cipher_TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,
	cipher_TLS_ECDH_RSA_WITH_AES_128_CBC_SHA256,
	cipher_TLS_ECDH_RSA_WITH_AES_256_CBC_SHA384,
	cipher_TLS_ECDH_ECDSA_WITH_AES_128_GCM_SHA256,
	cipher_TLS_ECDH_ECDSA_WITH_AES_256_GCM_SHA384,
	cipher_TLS_ECDH_RSA_WITH_AES_128_GCM_SHA256,
	cipher_TLS_ECDH_RSA_WITH_AES_256_GCM_SHA384,
	cipher_TLS_ECDHE_PSK_WITH_RC4_128_SHA,
	cipher_TLS_ECDHE_PSK_WITH_3DES_EDE_CBC_SHA,
	cipher_TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA,
	cipher_TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA,
	cipher_TLS_ECDHE_PSK_WITH_AES_128_CBC_SHA256,
	cipher_TLS_ECDHE_PSK_WITH_AES_256_CBC_SHA384,
	cipher_TLS_ECDHE_PSK_WITH_NULL_SHA,
	cipher_TLS_ECDHE_PSK_WITH_NULL_SHA256,
	cipher_TLS_ECDHE_PSK_WITH_NULL_SHA384,
	cipher_TLS_RSA_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_RSA_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_DH_DSS_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_DH_DSS_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_DH_RSA_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_DH_RSA_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_DHE_DSS_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_DHE_DSS_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_DHE_RSA_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_DHE_RSA_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_DH_anon_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_DH_anon_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_ECDHE_ECDSA_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_ECDHE_ECDSA_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_ECDHE_RSA_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_ECDHE_RSA_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_ECDH_RSA_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_ECDH_RSA_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_RSA_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_RSA_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_DH_RSA_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_DH_RSA_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_DH_DSS_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_DH_DSS_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_DH_anon_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_DH_anon_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_ECDH_ECDSA_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_ECDH_RSA_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_ECDH_RSA_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_PSK_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_PSK_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_DHE_PSK_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_DHE_PSK_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_RSA_PSK_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_RSA_PSK_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_PSK_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_PSK_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_RSA_PSK_WITH_ARIA_128_GCM_SHA256,
	cipher_TLS_RSA_PSK_WITH_ARIA_256_GCM_SHA384,
	cipher_TLS_ECDHE_PSK_WITH_ARIA_128_CBC_SHA256,
	cipher_TLS_ECDHE_PSK_WITH_ARIA_256_CBC_SHA384,
	cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_ECDHE_ECDSA_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_ECDHE_RSA_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_RSA_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_RSA_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_DH_RSA_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_DH_RSA_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_DH_DSS_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_DH_DSS_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_DH_anon_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_DH_anon_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_ECDH_ECDSA_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_ECDH_RSA_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_PSK_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_PSK_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_128_GCM_SHA256,
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_256_GCM_SHA384,
	cipher_TLS_PSK_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_PSK_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_DHE_PSK_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_DHE_PSK_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_RSA_PSK_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_ECDHE_PSK_WITH_CAMELLIA_128_CBC_SHA256,
	cipher_TLS_ECDHE_PSK_WITH_CAMELLIA_256_CBC_SHA384,
	cipher_TLS_RSA_WITH_AES_128_CCM,
	cipher_TLS_RSA_WITH_AES_256_CCM,
	cipher_TLS_RSA_WITH_AES_128_CCM_8,
	cipher_TLS_RSA_WITH_AES_256_CCM_8,
	cipher_TLS_PSK_WITH_AES_128_CCM,
	cipher_TLS_PSK_WITH_AES_256_CCM,
	cipher_TLS_PSK_WITH_AES_128_CCM_8,
	cipher_TLS_PSK_WITH_AES_256_CCM_8,
}
