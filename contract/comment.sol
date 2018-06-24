pragma solidity ^0.4.16;
/*
 * 保存留言在链上
 * @author:caige
 * @since: 2018-6-22
 */

contract comment {

    struct Comment{
        bytes name;
        bytes message;
        uint time;
    }

    Comment[] public commentList;

    function contructor(){

    }

    function set(bytes name,bytes message,uint time) public {
        commentList.push(Comment({name:name,message:message,time:time}));
    }

    function get(uint id) public view returns(bytes name,bytes message,uint time){
        Comment item = commentList[id];
        name = item.name;
        message = item.message;
        time = item.time;
    }

    function length() public view returns(uint len){
        len = commentList.length;
    }

}
