/*
* @Author: Ximidar
* @Date:   2018-09-02 01:37:29
* @Last Modified by:   Ximidar
* @Last Modified time: 2018-10-17 14:11:51
 */

package DataStructures

type ReplyString struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ReplyJSON struct {
	Success bool   `json:"success"`
	Message []byte `json:"message"`
}
