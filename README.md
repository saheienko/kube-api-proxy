# Demo time

```bash
➜  kube-api-proxy git:(master) ✗ cat <<EOF > examples/minikube.json
{
  "id": "",
  "apiHost": "$(minikube ip)",
  "apiPort": "8443",
  "auth": {
    "username": "minikube",
    "token": "",
    "ca": "$(while read -r line; do echo -n $line;echo -n '\\n'; done <~/.minikube/ca.crt)",
    "cert": "$(while read -r line; do echo -n $line;echo -n '\\n'; done <~/.minikube/client.crt)",
    "key": "$(while read -r line; do echo -n $line;echo -n '\\n'; done <~/.minikube/client.key)"
  }
}
EOF

➜  kube-api-proxy git:(master) ✗ curl -XPOST -d'@examples/minikube.json' localhost:8080/kubes
has been added: 66ba4df77dfeeb1b74d0c709ab7d866ad71396efa741c45ca48a548ef260ee82

➜  kube-api-proxy git:(master) ✗ curl -s localhost:8080/kubes/66ba4df77dfeeb1b74d0c709ab7d866ad71396efa741c45ca48a548ef260ee82 | jq
{
  "id": "66ba4df77dfeeb1b74d0c709ab7d866ad71396efa741c45ca48a548ef260ee82",
  "apiHost": "192.168.39.175",
  "apiPort": "8443",
  "auth": {
    "username": "minikube",
    "token": "",
    "ca": "-----BEGIN CERTIFICATE-----\nMIIC5zCCAc+gAwIBAgIBATANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p\na3ViZUNBMB4XDTE4MDYwNjE3MTkwMFoXDTI4MDYwMzE3MTkwMFowFTETMBEGA1UE\nAxMKbWluaWt1YmVDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALRE\n2qHp02w5lzZHjN0HOZxryli8W34GyPH6BpqzhcKJaUnuXbNSVBa1fXk/CpfgSP0h\nHZ4j1QxntDM8ZN8hzYeiOx+aAiy0jHO82w8wzo45+Jgi+mCIiYk9LDsXRrzieYZF\nwr1VNecduUoHTvOAQ1ExxPkQbe07xtwsL1gVHH0KGmzrYZtj1EZgEV8s17MrMumn\nE7RmvwhUu7itYdtQUo0WZ/r8HQUa3EIqCHqLV0T9nst4oEKHAhKSauj3uF8fx5UE\nmlaPcMSymLBbN0x8fhzKApP29qxwWqiexkhx/eZ/OEZRKap5TkuHvdMXM1D0grWl\ngLmhFn/5SoFPVDysdWcCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMB0GA1UdJQQW\nMBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3\nDQEBCwUAA4IBAQCZ+Rk/T/qqpNH8ipspks4LKoPANySUqsMYPK4y/ehnKQIfjVze\nkFruMmV0J8UkMEvS1lPNPquhpNNK7gnTQ6fMHGWXsqcp/XXvNmlQEurpjHd6mFY3\nKhUjp5mo3iRnzgci2B3/9xj17SBU0cIiVB4gU/05U6VnzmE973J85aK3J92bkLA9\nz7vK/OEzlMSZF0I2e2uFvvXWBNu5maSKlWRKkpEFZxzunYeFZZhnL3h6ahYA9WRp\nJm6WLAkUxLHXzLIco0k84xiHaIa+AHnn3pzl9nskFixbZ2ETtS5sv+wMuLqO0st/\nCIUJOQCVXMvkksafeqFe90pPZvbRBgUTvEbF\n-----END CERTIFICATE-----\n",
    "cert": "-----BEGIN CERTIFICATE-----\nMIIDADCCAeigAwIBAgIBAjANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p\na3ViZUNBMB4XDTE4MDYwNjE3MjcyNFoXDTE5MDYwNjE3MjcyNFowMTEXMBUGA1UE\nChMOc3lzdGVtOm1hc3RlcnMxFjAUBgNVBAMTDW1pbmlrdWJlLXVzZXIwggEiMA0G\nCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDSomUq69LNHutTeQXrqU7tE18z/I0x\nyEwakB5Uq1LSK2UIiIkLEXiuem/kr4fK+Dxk95O0NIqyp2uc7tn0FuBJFSd+7YHR\nXKH6mjCIyqKEQPGzj+I7fLGnz31nAVeboOOUv5/H9XdzNmo0DnRdz829puezJCzm\nsPxeQFc7EFeQbkRtQ1rts8AcQucEk7VlM2sp5EY2pk98cwZIQSBMnxg3wyQ2Cvdp\nOIUUvWVSNt6uRHeL9pkBdtbV2PwX6w+EmNima+Ufh3PwO6aLPBg7VGRelnSdUFCJ\n0umuPJJAlwGw5Y1O+NDKVQ8L08imlL6XAf6AXXdg5ey+aNb7yN96IeQxAgMBAAGj\nPzA9MA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUH\nAwIwDAYDVR0TAQH/BAIwADANBgkqhkiG9w0BAQsFAAOCAQEAcTxWQr1MhIcNfgEU\nfobH3viB4OxK2UwSENDclrvfoI40E9VNHhYrEKWR9OADGcSu5DS8iBleh+bUbNfQ\n6/j2eZ58nP7S26uM43ACk3plg3c4PyCXQhjxuAKejVMTMMfsZIdmKpes9jcRbPuI\n4Go2vAhJTGHvMhlJnpEbQ/FxYuml+3lfUTY6DRZtuTfd2R2pySA9INQKvjGkS2ZV\noAagmJWGW4MsI7NG5cDDGlaysa+5BCbk/rh+vdRFoUKudWkNGDm4rsLaZIqkt4wr\nKAw/oOD3MRIPMRpUcuW5EZ5wN91fkD6cOugWuOnWSWRJzZpL6xqvImOFneUnwv58\nLbe+fw==\n-----END CERTIFICATE-----\n",
    "key": "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEA0qJlKuvSzR7rU3kF66lO7RNfM/yNMchMGpAeVKtS0itlCIiJ\nCxF4rnpv5K+Hyvg8ZPeTtDSKsqdrnO7Z9BbgSRUnfu2B0Vyh+powiMqihEDxs4/i\nO3yxp899ZwFXm6DjlL+fx/V3czZqNA50Xc/NvabnsyQs5rD8XkBXOxBXkG5EbUNa\n7bPAHELnBJO1ZTNrKeRGNqZPfHMGSEEgTJ8YN8MkNgr3aTiFFL1lUjberkR3i/aZ\nAXbW1dj8F+sPhJjYpmvlH4dz8DumizwYO1RkXpZ0nVBQidLprjySQJcBsOWNTvjQ\nylUPC9PIppS+lwH+gF13YOXsvmjW+8jfeiHkMQIDAQABAoIBACQ2PRRS9KvFDAoO\nvWDVe7cwZGaonZGYcNUEP+KojZWKVlVQO9dGSqwcao4zSzIu2Rs2oRMTEWFDfTG+\nsoPPRwHpfB/LL01SEprl1UA/Lg90ptkK/IbjmhtShammxmwADgAtrYeQANgy27FV\nZtYV+rYHMsBOkNWcSdbeUuDZn2Q58LjqDS39t0ivHxi9pt2sfi0uOKmZbT5ecYht\ngK4+Feom0vyjR8ohDF7LZcCdLMdwwhMS3VjyNYRQBckXZyDUWSeg4avz3SE15Ofb\nkaDhFqG9Rpnh+9Wn0lOHGeJBivUVP+xYxYkBf2BFtZ8HBlCgYTp1Aw5QszsuIK+j\nDb/fPcECgYEA+7vulu5mnNvC5qdJ2mGQUyJwdNFqXE68GupPXa4C/z3slFiR8bpW\nUP52pPgPCM2qgNNXNrd7mVwmIqTf5TeCMT2eMD5RKT5jjh7Sp9zGvoTVBx0zuVNr\nfGRwn4vYm0XF27GdzKhPw2LGYISIuXLE8fHFWU5D98tCwI0N5co8pMkCgYEA1jQq\nQIgEbG2MBqN65ZhLO/tDyfJr/HQW2cLNlej1erHLif4YJuRrIgb3sRsg61W9Zaw3\n20LVlJ9EbAtLUzGtBfnDnIyQOHKn5qtuBIBkQso9NIjkOcqCPQU8rnGAgKuC04vv\nbxBjZe5rFrqVIocscXFkWMy5mOLZeYZvPJpPgCkCgYEAosFK7QKODXR4erBGK49Q\nxK9LjfunjK7LJ4u+bI8JGQVsZC0vjt4u2IbtJpPLBKIUTt5VUOcoXmsZrOR0bbqJ\nzlRMZlykFMpli4maITW4uY0gPk0/F987a111A3JjRWDDH9uibqOTjnvaTqTh0STG\n+LacJbVYdGlSazPHfH5Y3yECgYA0JQrMHtCE3L4jt5RpZAOcnHRKKxuin1gYttV5\nUva/YZzdAOA8R4rVA8E0ehgvcfXjVGNcmw6HWaY8bxttK0ClncHC0G0jcLXy73Se\n3+qIX9c6fMCiWOwPksDM7pCLwjTc7sngzaqE299x7wXzG9jz3NjCzUO5NjAe510Y\n8a+80QKBgC5FLEQokSu1BQ9YEq5Xm03iEKvcm/TkFRFheGxNvzw5+yb7Z4KGAmDo\nog7swAR1QiiMfigrOsSx1J64ePJn/B5Jcygvkyz/JWqUUIC6CZf5J2NiIlyfNAeg\n2UGKk5VaQ2SyfnTGFVCf0epyyhvRxRCnvtPgiHcJC0iqw7st1oD1\n-----END RSA PRIVATE KEY-----\n"
  }
}

➜  kube-api-proxy git:(master) ✗ curl -s localhost:8080/kubes/66ba4df77dfeeb1b74d0c709ab7d866ad71396efa741c45ca48a548ef260ee82/resources/namespaces | jq
{
  "kind": "NamespaceList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/namespaces",
    "resourceVersion": "11410"
  },
  "items": [
    {
      "metadata": {
        "name": "default",
        "selfLink": "/api/v1/namespaces/default",
        "uid": "f6b1c459-69ae-11e8-a279-eccf47fa16cb",
        "resourceVersion": "12",
        "creationTimestamp": "2018-06-06T17:28:01Z"
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    },
    {
      "metadata": {
        "name": "kube-public",
        "selfLink": "/api/v1/namespaces/kube-public",
        "uid": "f939b5f1-69ae-11e8-a279-eccf47fa16cb",
        "resourceVersion": "45",
        "creationTimestamp": "2018-06-06T17:28:06Z"
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    },
    {
      "metadata": {
        "name": "kube-system",
        "selfLink": "/api/v1/namespaces/kube-system",
        "uid": "f6961feb-69ae-11e8-a279-eccf47fa16cb",
        "resourceVersion": "15",
        "creationTimestamp": "2018-06-06T17:28:01Z",
        "annotations": {
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"                                                                                                                                                              
        }
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    }
  ]
}

➜  kube-api-proxy git:(master) ✗ kubectl get ns -o json                                                                                                 
{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "v1",
            "kind": "Namespace",
            "metadata": {
                "creationTimestamp": "2018-06-06T17:28:01Z",
                "name": "default",
                "namespace": "",
                "resourceVersion": "12",
                "selfLink": "/api/v1/namespaces/default",
                "uid": "f6b1c459-69ae-11e8-a279-eccf47fa16cb"
            },
            "spec": {
                "finalizers": [
                    "kubernetes"
                ]
            },
            "status": {
                "phase": "Active"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Namespace",
            "metadata": {
                "creationTimestamp": "2018-06-06T17:28:06Z",
                "name": "kube-public",
                "namespace": "",
                "resourceVersion": "45",
                "selfLink": "/api/v1/namespaces/kube-public",
                "uid": "f939b5f1-69ae-11e8-a279-eccf47fa16cb"
            },
            "spec": {
                "finalizers": [
                    "kubernetes"
                ]
            },
            "status": {
                "phase": "Active"
            }
        },
        {
            "apiVersion": "v1",
            "kind": "Namespace",
            "metadata": {
                "annotations": {
                    "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Namespace\",\"metadata\":{\"annotations\":{},\"name\":\"kube-system\",\"namespace\":\"\"}}\n"
                },
                "creationTimestamp": "2018-06-06T17:28:01Z",
                "name": "kube-system",
                "namespace": "",
                "resourceVersion": "15",
                "selfLink": "/api/v1/namespaces/kube-system",
                "uid": "f6961feb-69ae-11e8-a279-eccf47fa16cb"
            },
            "spec": {
                "finalizers": [
                    "kubernetes"
                ]
            },
            "status": {
                "phase": "Active"
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": "",
        "selfLink": ""
    }
}

➜  kube-api-proxy git:(master) ✗ curl -s "localhost:8080/kubes/66ba4df77dfeeb1b74d0c709ab7d866ad71396efa741c45ca48a548ef260ee82/resources/services?namespace=default" | jq 
{
  "kind": "ServiceList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/namespaces/default/services",
    "resourceVersion": "11533"
  },
  "items": [
    {
      "metadata": {
        "name": "kubernetes",
        "namespace": "default",
        "selfLink": "/api/v1/namespaces/default/services/kubernetes",
        "uid": "f926d180-69ae-11e8-a279-eccf47fa16cb",
        "resourceVersion": "39",
        "creationTimestamp": "2018-06-06T17:28:06Z",
        "labels": {
          "component": "apiserver",
          "provider": "kubernetes"
        }
      },
      "spec": {
        "ports": [
          {
            "name": "https",
            "protocol": "TCP",
            "port": 443,
            "targetPort": 8443
          }
        ],
        "clusterIP": "10.96.0.1",
        "type": "ClusterIP",
        "sessionAffinity": "ClientIP",
        "sessionAffinityConfig": {
          "clientIP": {
            "timeoutSeconds": 10800
          }
        }
      },
      "status": {
        "loadBalancer": {}
      }
    }
  ]
}

➜  kube-api-proxy git:(master) ✗ kubectl get service --namespace=default -o json
{
    "apiVersion": "v1",
    "items": [
        {
            "apiVersion": "v1",
            "kind": "Service",
            "metadata": {
                "creationTimestamp": "2018-06-06T17:28:06Z",
                "labels": {
                    "component": "apiserver",
                    "provider": "kubernetes"
                },
                "name": "kubernetes",
                "namespace": "default",
                "resourceVersion": "39",
                "selfLink": "/api/v1/namespaces/default/services/kubernetes",
                "uid": "f926d180-69ae-11e8-a279-eccf47fa16cb"
            },
            "spec": {
                "clusterIP": "10.96.0.1",
                "ports": [
                    {
                        "name": "https",
                        "port": 443,
                        "protocol": "TCP",
                        "targetPort": 8443
                    }
                ],
                "sessionAffinity": "ClientIP",
                "sessionAffinityConfig": {
                    "clientIP": {
                        "timeoutSeconds": 10800
                    }
                },
                "type": "ClusterIP"
            },
            "status": {
                "loadBalancer": {}
            }
        }
    ],
    "kind": "List",
    "metadata": {
        "resourceVersion": "",
        "selfLink": ""
    }
}
```

