package test

// SSLTestPath is the path intermediary SSL certificates are created in for test cases
const SSLTestPath = "ssl"

// CACertFolder is the folder at which intermediary CA certificates are created in for test cases
const CACertFolder = "ssl/ca"

// CACertPath is the path of the intermediary CA certificate
const CACertPath = "ssl/ca/ca_cert.pem"

// RegistrationCert is a sample transaction relay registration cert
const RegistrationCert = "-----BEGIN CERTIFICATE-----\nMIICtjCCAjygAwIBAgIUW8vj+TuK43zqfd7Fj9xNqkf0Fa8wCgYIKoZIzj0EAwIw\nbzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFu\nc3RvbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMR0wGwYDVQQDDBRibG9Y\ncm91dGUudGVzdG5ldC5DQTAeFw0yMTAyMDgxNzUzMDNaFw0yMjAyMDgwMDAwMDBa\nMGwxCzAJBgNVBAYTAiAgMQswCQYDVQQIDAIgIDELMAkGA1UEBwwCICAxGzAZBgNV\nBAoMEmJsb1hyb3V0ZSBMQUJTIEluYzEmMCQGA1UEAwwdYmxvWHJvdXRlLnRlc3Ru\nZXQucmVsYXlfcHJveHkwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAAQ9HXj7YMkHI+7P\nxh5uarJoEmOF2s2/PT/+9Vu4rAdMUZQ0e5jELdTEuo/JGT1AlFwquhiSPAsVroEX\nseZyfViOSajhUBmTYz+LljKilltGzK8Lkbw06jQa08NWX1bMwx2jgZswgZgwHwYD\nVR0jBBgwFoAUkN1jVx3ISR8ATM29MLuIZ0K0dJEwKAYDVR0RBCEwH4IdYmxvWHJv\ndXRlLnRlc3RuZXQucmVsYXlfcHJveHkwDAYDVR0TAQH/BAIwADAOBgNVHQ8BAf8E\nBAMCBeAwFAYFPoJNolwEC1JFTEFZX1BST1hZMBcGBT6CTaJeBA5ibG9Ycm91dGUg\nTEFCUzAKBggqhkjOPQQDAgNoADBlAjAUrqyAz8gun5mI02E0MtiqcH3aVNayIvyH\nc7/c1+qyKLobJ0DwKqaocX4iYZrs7e0CMQC3wRz/P2bIxVeViE8/nnRQKLekBQqp\nBTwJ9D/G46Tg7r6kxj9QGQD8onp/i5AUCOI=\n-----END CERTIFICATE-----\n"

// RegistrationKey is a sample transaction relay registration key
const RegistrationKey = "-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDBX1DtibI/7qGv2jkw0/uSm5nF2OfsGWuXiNcz5Sy3RlsvvKohqb0U2\nVjdqSG9KnESgBwYFK4EEACKhZANiAAQ9HXj7YMkHI+7Pxh5uarJoEmOF2s2/PT/+\n9Vu4rAdMUZQ0e5jELdTEuo/JGT1AlFwquhiSPAsVroEXseZyfViOSajhUBmTYz+L\nljKilltGzK8Lkbw06jQa08NWX1bMwx0=\n-----END EC PRIVATE KEY-----\n"

// PrivateCert is a sample transaction relay private cert
const PrivateCert = "-----BEGIN CERTIFICATE-----\nMIIDCTCCAo+gAwIBAgIUBPo2XvfCGvlFy6Cl0LPXej+Q5wQwCgYIKoZIzj0EAwIw\nbzELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFu\nc3RvbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMR0wGwYDVQQDDBRibG9Y\ncm91dGUudGVzdG5ldC5DQTAeFw0yMTAyMDgyMDAyMDNaFw0yMjAyMDgwMDAwMDBa\nMHIxCzAJBgNVBAYTAiAgMQswCQYDVQQIDAIgIDELMAkGA1UEBwwCICAxGzAZBgNV\nBAoMEmJsb1hyb3V0ZSBMQUJTIEluYzEsMCoGA1UEAwwjYmxvWHJvdXRlLnRlc3Ru\nZXQucmVsYXlfdHJhbnNhY3Rpb24wdjAQBgcqhkjOPQIBBgUrgQQAIgNiAASMommd\nDnH8xgfXKM0HRln9vTk7O0Y1bZp+oRdq3Mv0xEY2lQJnWNTL8QqpfcFlUgJDiCRB\nP9SlS2H/DIbGah5yWv1XIknGkHXaS9o36PIYDAWG7UOwMinSpZ77sdOzq9Wjgegw\ngeUwHwYDVR0jBBgwFoAUkN1jVx3ISR8ATM29MLuIZ0K0dJEwLgYDVR0RBCcwJYIj\nYmxvWHJvdXRlLnRlc3RuZXQucmVsYXlfdHJhbnNhY3Rpb24wDAYDVR0TAQH/BAIw\nADAOBgNVHQ8BAf8EBAMCBeAwGgYFPoJNolwEEVJFTEFZX1RSQU5TQUNUSU9OMC0G\nBT6CTaJdBCQwZjU0YzUwOS0wNmYwLTRiZGQtOGZjMC0zYmRmMWFjMTE5ZWQwFwYF\nPoJNol4EDmJsb1hyb3V0ZSBMQUJTMBAGBT6CTaJfBAdnZW5lcmFsMAoGCCqGSM49\nBAMCA2gAMGUCMHbPKeumEf46NN+FSE9rKCD6bHH9pcm1asaF5XfSdiSzaiOseus+\njIqGOY0L8tKnPwIxAO9XUyoOzutvjrqIJ+APi6imUNKp3GaQeRMdXDXJH4vHoJ6L\nxpn1wK0JgqdzdA2z8w==\n-----END CERTIFICATE-----\n"

// PrivateKey is a sample transaction relay private key
const PrivateKey = "-----BEGIN EC PRIVATE KEY-----\nMIGkAgEBBDBEU5Aui2ua3hSaWwo1YCH+DQAdu0aQl5N67GS8ZeQr0XnucDasQH0f\n84FLiQ6m7SqgBwYFK4EEACKhZANiAASMommdDnH8xgfXKM0HRln9vTk7O0Y1bZp+\noRdq3Mv0xEY2lQJnWNTL8QqpfcFlUgJDiCRBP9SlS2H/DIbGah5yWv1XIknGkHXa\nS9o36PIYDAWG7UOwMinSpZ77sdOzq9U=\n-----END EC PRIVATE KEY-----\n"

// CACert is a sample CA certificate
const CACert = "-----BEGIN CERTIFICATE-----\nMIICkDCCAhWgAwIBAgIUQGhfIhpMxaSE0r/jjB35VclkTYgwCgYIKoZIzj0EAwIw\nazELMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFu\nc3RvbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMRkwFwYDVQQDDBBibG9Y\ncm91dGUuZGV2LkNBMB4XDTIwMTEwOTIyMzY0NloXDTQ4MDMyNzAwMDAwMFowazEL\nMAkGA1UEBhMCVVMxETAPBgNVBAgMCElsbGlub2lzMREwDwYDVQQHDAhFdmFuc3Rv\nbjEbMBkGA1UECgwSYmxvWHJvdXRlIExBQlMgSW5jMRkwFwYDVQQDDBBibG9Ycm91\ndGUuZGV2LkNBMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEd5IkD4wqWVGbq0jCehjr\nyvOEkbD5vCYIes4UwRH9Z7YSfeKQrOSaW1LUzHCYvqOTLfgHoAC0lS1v9+OlT/LZ\nboey8h4TZwLZ9zLbHzyVcTww01ZFNeBndQ2EmdaSYdqKo3oweDAfBgNVHSMEGDAW\ngBQHnCDGbOz7WL0VB9Kd0YdQSpsnQzAdBgNVHQ4EFgQUB5wgxmzs+1i9FQfSndGH\nUEqbJ0MwGwYDVR0RBBQwEoIQYmxvWHJvdXRlLmRldi5DQTAMBgNVHRMEBTADAQH/\nMAsGA1UdDwQEAwIB/jAKBggqhkjOPQQDAgNpADBmAjEAo6kPPChOytP961lFjKFb\n+zfEPm6sHtBxmgeDMhQwqb1erIIsYfU6zVaA82g9REHvAjEAoLfzcjEq91/Jlcmn\nCSgJY3JUPIocBek+o9cKczwz1ZDuzGscMOF0J4fpTwAyJOUP\n-----END CERTIFICATE-----"
