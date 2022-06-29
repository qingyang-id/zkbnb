/*
 * Copyright © 2021 Zecrey Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package util

import (
	"math/big"

	"github.com/zecrey-labs/zecrey-crypto/ffmath"
	"github.com/zecrey-labs/zecrey-crypto/util"
	"github.com/zecrey-labs/zecrey-legend/pkg/zerror"
)

/*
	ToPackedAmount: convert big int to 40 bit, 5 bits for 10^x, 35 bits for a * 10^x
*/
func ToPackedAmount(amount *big.Int) (res int64, err error) {
	return util.ToPackedAmount(amount)
}

func CleanPackedAmount(amount *big.Int) (nAmount *big.Int, err error) {
	return util.CleanPackedAmount(amount)
}

/*
	ToPackedFee: convert big int to 16 bit, 5 bits for 10^x, 11 bits for a * 10^x
*/
func ToPackedFee(amount *big.Int) (res int64, err error) {
	return util.ToPackedFee(amount)
}

func CleanPackedFee(amount *big.Int) (nAmount *big.Int, err error) {
	return util.CleanPackedFee(amount)
}

func CheckPackedFee(gassFee *big.Int) (err error) {
	checkParm, err := util.CleanPackedFee(gassFee)
	if err != nil {
		return err
	}
	if !ffmath.Equal(gassFee, checkParm) {
		return zerror.New(20002, "Invalid Gas Fee")
	}
	return nil
}

func CheckPackedAmount(amount *big.Int) (err error) {
	checkParm, err := util.CleanPackedAmount(amount)
	if err != nil {
		return err
	}
	if !ffmath.Equal(amount, checkParm) {
		return zerror.New(20003, "Invalid amount")
	}
	return nil
}
