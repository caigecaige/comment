pragma solidity ^0.4.16;
/*
 * 保存留言在链上
 * @author:caige
 * @since: 2018-6-22
 */

contract comment {

    //留言的struct
    struct Comment{
        bytes name;
        bytes message;
        uint time;
    }

    //保存所有的留言
    Comment[] public commentList;

    function contructor(){

    }

    //保存留言信息
    function set(bytes name,bytes message,uint time) public {
        commentList.push(Comment({name:name,message:message,time:time}));
    }

    //根据index获取单个留言信息
    function get(uint id) public view returns(bytes name,bytes message,uint time){
        Comment item = commentList[id];
        name = item.name;
        message = item.message;
        time = item.time;
    }

    //留言数组的长度
    function length() public view returns(uint len){
        len = commentList.length;
    }

}
