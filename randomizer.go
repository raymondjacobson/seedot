/* randomizing functions */
package main

import (
  "math/rand"
  "time"
)

/*returns a random string of between minsize and minsize + variance */
func RandomString(minsize, variance int) string {
  abet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890!@$^&*()-=[];.,<>:}{}'"
  random_string := ""
  rand.Seed(time.Now().UnixNano())
  for i:=0; i<minsize+rand.Intn(variance); i++ {
    random_string+=string(abet[rand.Intn(len(abet))])
  }
  return random_string
}
