一个还行的中文文档：https://learnblockchain.cn/docs/hardhat/getting-started/
一个还行的入门视频：https://www.bilibili.com/video/BV1RFsfe5Ek5

# 安装hardhat

1. 用`npm`工具初始化项目：`npm init`，默认会以当前文件夹名来作为项目名。完成之后，会生成一个`package.json`文件
2. 安装`hardhat`：`npm install hardhat --save-dev`，完成之后，会出现一个`package-lock.json`文件
3. 初始化项目：`npx hardhat`，完成之后，会出现一个`hardhat.config.js`文件

# 配置依赖

安装`npm install @chainlink/contracts --save-dev`

# 部署合约

通过脚本部署：创建一个文件`deployFundMe.js`，在里面编写部署的命令。然后执行语句`npx hardhat run scripts/deployFundMe.js --network sepolia`，如果不指定network，默认是`hardhat`本地网络。部署成功之后会返回合约的地址。

要想通过配置文件来存储一些配置信息的话，需要引用dotenv这个工程，`npm install dotenv --save-dev`，然后配置.env文件。接着在`hardhat.config.js`文件中，引入`dotenv`，`require("dotenv").config();`，配置参数即可。种方式仍然是通过明文存储的，不太安全。

可以通过`chainlink`提供的`env-enc`来处理，`npm install @chainlink/env-enc --save-dev`。然后执行`npx env-enc set-pw`来设置一个密码，这里的密码是123456。然后就可以设置各种各样的环境变量了。先执行`npx env-enc set`，然后输入变量名，然后输入变量值。接着就会看到生成了一个`.env.enc`文件，里面存储了加密后的变量信息。配置完之后在`hardhat.config.js`文件中引入`env-enc`即可，`require("@chainlink/env-enc").config();`。

**配置文件在上传的时候要万分小心！！！**

# 验证合约

`npx hardhat verify --network sepolia 合约地址 构造函数参数`。这里的构造参数可以是deploy合约时传入的参数，也可以是空。在`deployFundMe.js`文件中，参数就是10。

>  `await fundMeFactory.deploy(10)`

这里需要配置API_KEY，见`hardhat.config.js`文件。如果验证一直报错，可以在机场里打开tun模式。网络问题，要使用局域网代理，搜一下github有一个解决文章。

也可以在`deployFundMe.js`文件中，配置验证的代码。

# tasks

hardhat中的所有命令，比如`run`都是一个task，用户也可以自己定义task。首先，创建一个路径tasks，然后在其中编写对应的js文件，比如`deploy.js`。然后，在`hardhat.config.js`文件中引入，然后就可以使用这个task了。配置完之后，就可以在`npx hardhat help`中看到这些task了。

# 测试

hardhat提供了一些测试框架，比如`mocha`和`chai`。只需要在test文件夹下写对应的测试js文件，然后执行`npx hardhat test`即可。

# hardhat-deploy

`npm install --save-dev hardhat-deploy`。然后在deploy文件夹下写对应的js文件，在`hardhat.config.js`引入，然后执行`npx hardhat deploy`即可。

# hardhat-ethers

`npm install --save-dev  @nomiclabs/hardhat-ethers hardhat-deploy-ethers ethers`

# hardhat-gas-reporter

`npm install --save-dev hardhat-gas-reporter`，用于监测合约的gas使用情况。