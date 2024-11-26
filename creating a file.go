// ----- FILE IO -----
	// We can create, write and read from files

	// Create a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Says to close the file after program ends or when
	// there is a closing curly bracket
	defer f.Close()

	// Create list of primes
	iPrimeArr := []int{2, 3, 5, 7, 11}
	// Create string array from int array
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	// Cycle through strings and write to file
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the created file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read from file and print once per line
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
