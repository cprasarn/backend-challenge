package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type InvalidEncodedString struct {
}

func (e *InvalidEncodedString) Error() string {
	return "Invalid encoded string"
}

func calc(s string) int {
	sv, err := build(s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("SV:", sv)

	l := len(sv)
	result := 0
	for i := 0; i < l; i++ {
		n, err := strconv.Atoi(string(sv[i]))
		if err != nil {
			break
		}
		result += n
	}
	return result
}

func build(s string) (string, error) {
	result := ""
	sl := len(s)
	l := 0
	r := 0
	e := 0
	for i := 0; i < sl; i++ {
		if rune(s[i]) == 76 {
			l++
			if r > 0 {
				rl := len(result)
				for j := 0; j < r; j++ {
					if rl == 0 {
						result += strconv.Itoa(j)
						result += strconv.Itoa(j + 1)
					} else {
						last, err := strconv.Atoi(string(result[rl-1]))
						if err != nil {
							break
						}
						last++
						if last > 9 {
							return "", &InvalidEncodedString{}
						}
						result += strconv.Itoa(last)
					}
					rl = len(result)
				}
				r = 0
			}
		}

		if rune(s[i]) == 82 {
			r++
			if l > 0 {
				rl := len(result)
				if rl > 0 {
					last, err := strconv.Atoi(string(result[rl-1]))
					if err != nil {
						log.Fatalln(err)
					}
					if last-l < 0 {
						if rl == 1 {
							result = strconv.Itoa(l)
						} else {
							eq := 0
							for i := rl - 1; i > 0; i-- {
								if rune(result[i]) == rune(result[i-1]) {
									eq++
								} else {
									break
								}
							}

							if eq == 0 && rune(result[rl-1]) > rune(result[rl-2]) {
								result = result[:rl-1] + strconv.Itoa(l)
							} else {
								result = result[:rl-eq-1] + strconv.Itoa(l)
								for i := eq; i > 0; i-- {
									result += strconv.Itoa(l)
								}
							}
						}
					}
				}

				for j := l; j > 0; j-- {
					if rl == 0 {
						result += strconv.Itoa(j)
						result += strconv.Itoa(j - 1)
					} else {
						rl = len(result)
						last, err := strconv.Atoi(string(result[rl-1]))
						if err != nil {
							break
						}
						last--
						if last >= 0 {
							result += strconv.Itoa(last)
						}
					}
					rl = len(result)
				}
				l = 0
			}
		}

		if rune(s[i]) == 61 {
			e++
			if r > 0 {
				rl := len(result)
				for j := 0; j < r; j++ {
					if rl == 0 {
						result += strconv.Itoa(j)
						result += strconv.Itoa(j + 1)
					} else {
						last, err := strconv.Atoi(string(result[rl-1]))
						if err != nil {
							break
						}
						last++
						if last > 9 {
							return "", &InvalidEncodedString{}
						}

						result += strconv.Itoa(last)
					}
					rl = len(result)
				}
				r = 0
			}

			if l > 0 && l < 10 {
				rl := len(result)
				if rl > 0 {
					last, err := strconv.Atoi(string(result[rl-1]))
					if err != nil {
						log.Fatalln(err)
					}
					if last-l < 0 {
						if rl == 1 {
							result = strconv.Itoa(l)
						} else {
							eq := 0
							for i := rl - 1; i > 0; i-- {
								if rune(result[i]) == rune(result[i-1]) {
									eq++
								} else {
									break
								}
							}

							if eq == 0 && rune(result[rl-1]) > rune(result[rl-2]) {
								result = result[:rl-1] + strconv.Itoa(l)
							} else {
								result = result[:rl-eq-1] + strconv.Itoa(l)
								for i := eq; i > 0; i-- {
									result += strconv.Itoa(l)
								}
							}
						}
					}
				}

				for j := l; j > 0; j-- {
					if rl == 0 {
						result += strconv.Itoa(j)
						result += strconv.Itoa(j - 1)
					} else {
						result += strconv.Itoa(j - 1)
					}
					rl = len(result)
				}
				l = 0
			} else if l > 9 {
				return "", &InvalidEncodedString{}
			}

			rl := len(result)
			if rl == 0 {
				result += strconv.Itoa(0)
				result += strconv.Itoa(0)
			} else {
				result += string(result[rl-1])
			}
			e = 0
		}
	}

	if e > 0 {
		rl := len(result)
		for i := 0; i < e; i++ {
			if rl == 0 {
				result += strconv.Itoa(0)
				result += strconv.Itoa(0)
			} else {
				result += string(result[rl-1])
			}
			rl = len(result)
		}
	}

	if l > 0 && l < 10 {
		rl := len(result)
		if rl > 0 {
			last, err := strconv.Atoi(string(result[rl-1]))
			if err != nil {
				log.Fatalln(err)
			}
			if last-l < 0 {
				if rl == 1 {
					result = strconv.Itoa(l)
				} else {
					eq := 0
					for i := rl - 1; i > 0; i-- {
						if rune(result[i]) == rune(result[i-1]) {
							eq++
						} else {
							break
						}
					}

					if eq == 0 && rune(result[rl-1]) > rune(result[rl-2]) {
						result = result[:rl-1] + strconv.Itoa(l)
					} else {
						result = result[:rl-eq-1] + strconv.Itoa(l)
						for i := eq; i > 0; i-- {
							result += strconv.Itoa(l)
						}
					}
				}
			}
		}

		for i := l; i > 0; i-- {
			if rl == 0 {
				result += strconv.Itoa(i)
				result += strconv.Itoa(i - 1)
			} else {
				last, err := strconv.Atoi(string(result[rl-1]))
				if err != nil {
					break
				}
				last--
				if last >= 0 {
					result += strconv.Itoa(last)
				}
			}
			rl = len(result)
		}
	} else if l > 9 {
		return "", &InvalidEncodedString{}
	}

	if r > 0 && r < 10 {
		rl := len(result)
		for i := 0; i < r; i++ {
			if rl == 0 {
				result += strconv.Itoa(i)
				result += strconv.Itoa(i + 1)
			} else {
				last, err := strconv.Atoi(string(result[rl-1]))
				if err != nil {
					return "", err
				}
				last++
				if last > 9 {
					return "", &InvalidEncodedString{}
				}
				result += strconv.Itoa(last)
			}
			rl = len(result)
		}
	} else if r > 9 {
		return "", &InvalidEncodedString{}
	}

	return result, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<encoded string>")
		os.Exit(1)
	}

	input := os.Args[1]
	fmt.Println("Input:", input)

	value := calc(input)
	fmt.Println("Output:", value)
}
