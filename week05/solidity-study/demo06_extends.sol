// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoExtendsGrandfather {
    
    event Log(string msg);

    modifier exactDividedBy2And3(uint256 _a) virtual {
        require(_a%2==0 && _a%3==0);
        _;
    }

    function hip() external virtual {
        emit Log("Grandfather msg");
    }

    function pop() external virtual {
        emit Log("Grandfather msg");
    }
}

// 简单的单继承
contract DemoExtendsFather is DemoExtendsGrandfather {

    // 重写合约
    modifier exactDividedBy2And3(uint _a) override virtual  {
        _;
        require(_a % 2 == 0 && _a % 3 == 0);
    }

    function hip() external override virtual {
        emit Log("Father msg");
    }

    function pip() external virtual {
        emit Log("Father msg");
    }

    function fatherOnly() public {
        emit Log("Father only msg");
    }
}

// 多重继承，继承必须按照辈分从高到低
contract DemoExtendsSon is DemoExtendsGrandfather, DemoExtendsFather {

    modifier exactDividedBy2And3(uint _a) override(DemoExtendsGrandfather, DemoExtendsFather) {
        _;
        require(_a % 2 == 0 && _a % 3 == 0);
    }

    // 当继承的多个合约都声明了同一个方法时，override中必须显示声明这些父合约
    function hip() external override(DemoExtendsGrandfather, DemoExtendsFather) {
        emit Log("Son msg");
    }

    // 调用父合约方法
    function callFatherHip() external {
        DemoExtendsFather.fatherOnly();
    }

    // 调用最近的（从右到左）父合约的方法
    function callSuperHip() external {
        super.fatherOnly();
    }
}

contract Dao {
    uint256 num;

    constructor (uint256 _num) {
        num = _num;
    }
}

// 构造器继承
contract MySqlDao is Dao(1) {}

contract PostgresDao is Dao{
    // 构造器继承
    constructor(uint256 _num) Dao(_num+1) {}
}

contract God {
    event Log(string msg);

    function foo() public virtual {
        emit Log("God.foo called");
    }

    function bar() public virtual {
        emit Log("God.bar called");
    }
}

contract Adam is God {
    function foo() public override(God) virtual {
        emit Log("Adam.foo called");
        super.foo();
    }

    function bar() public override(God) virtual {
        emit Log("Adam.bar called");
        super.bar();
    }
}

contract Eva is God {
    function foo() public override(God) virtual {
        emit Log("Eva.foo called");
        super.foo();
    }

    function bar() public override(God) virtual {
        emit Log("Eva.bar called");
        super.bar();
    }
}

// 菱形继承
contract People is Adam, Eva {
    function foo() public override(Adam, Eva) virtual {
        // 这里会从右到左，依次调用foo()，但是最后只会调用一次God.foo()
        // Eva.foo() -> Adam.foo() -> God.foo()
        // 原因是Solidity借鉴了Python的方式，强制一个由基类构成的DAG（有向无环图）使其保证一个特定的顺序
        super.foo();
    }

    function bar() public override(Adam, Eva) virtual {
        super.bar();
    }
}