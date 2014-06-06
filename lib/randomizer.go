/* randomizing functions */

func randomString(minsize, variance int) string {
  abet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890!@#$%^&*()-=[];.,/<>?:}{}'"
  random_string := ""
  rand.Seed(time.Now().UnixNano())
  for i:=0; i<minsize+rand.Intn(variance); i++ {
    random_string+=abet[rand.Intn(len(abet))]
  }
  return random_string
}
