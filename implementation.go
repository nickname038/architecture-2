package lab2

import (
	"errors"
	"regexp"
	"strings"
)

type symbolObject struct{
	value string
	isModified bool
}

func isOperand(obj symbolObject) (result bool, err error) {
	numbersPattern := `^\d+$`
	isObjValueNumber, err := regexp.MatchString(numbersPattern, obj.value)
	return (isObjValueNumber || obj.isModified), err
}

func isOperator(obj symbolObject) (result bool) {
	str := obj.value
  return str == "+" || str == "-" || str == "*" || str == "/" || str == "^"
}


// This function convert string with expression in postfix form
// to prefix form.
// If input string is invalid return error
func PostfixToPrefix(input string) (result string, err error) {
	defer func() {
    if recover() != nil {
      err = errors.New("invalid expression")
    }
  }()

  input = strings.Trim(input, " ")
	symbols := strings.Split(input, " ")

  if len(symbols) < 3 {
    panic("invalid expression")
  }

	objects := []symbolObject{}

	for i := 0; i < len(symbols); i++ {
		newObject := symbolObject {symbols[i], false}
		objects = append(objects, newObject)
	}

	isAllObjectsChecked := false
	for len(objects) > 1 && !isAllObjectsChecked {
		for i := 0; i < len(symbols) - 1; i++ {
			firstElem := objects[i]
			secondElem := objects[i + 1]
			thirdElem := objects[i + 2]

			isFirstElemOperand, err1 := isOperand(firstElem)
			isSecondElemOperand, err2 := isOperand(secondElem)
			isThirdElemOperator := isOperator(thirdElem)

			if err1 != nil || err2 != nil {
				err = errors.New("error")
				return
			}

			if isFirstElemOperand && isSecondElemOperand && isThirdElemOperator {
				firstPartOfNewSlice := objects[:i]
				objInMiddleOfNewSlice := symbolObject {thirdElem.value + " " + firstElem.value + " " + secondElem.value, true}
        thirdPartOfNewSlice := objects[i + 3:]

				objects = append(firstPartOfNewSlice, objInMiddleOfNewSlice)
				objects = append(objects, thirdPartOfNewSlice...)

				break
			}

			if (i + 2 == len(objects) - 1) {
				isAllObjectsChecked = true
			}
		}
	}

  result = objects[0].value
	return
}
