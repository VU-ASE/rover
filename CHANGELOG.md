# Changelog

## [1.1.6](https://github.com/VU-ASE/rover/compare/v1.1.5...v1.1.6) (2025-05-21)


### Bug Fixes

* separated upload and added debug prints ([37718dd](https://github.com/VU-ASE/rover/commit/37718dd825fed1fd0c88aa6534bc3749afe5bb47))

## [1.1.5](https://github.com/VU-ASE/rover/compare/v1.1.4...v1.1.5) (2025-05-21)


### Bug Fixes

* testing upload CI script ([1a73412](https://github.com/VU-ASE/rover/commit/1a7341229eaeb66e1884195412a7f86492e40181))

## [1.1.4](https://github.com/VU-ASE/rover/compare/v1.1.3...v1.1.4) (2025-05-21)


### Bug Fixes

* typo in makefile :( ([f581038](https://github.com/VU-ASE/rover/commit/f5810380aefa2134b918ce054252062863a5eb44))

## [1.1.3](https://github.com/VU-ASE/rover/compare/v1.1.2...v1.1.3) (2025-05-21)


### Bug Fixes

* removed dev setup script from build target ([369c1d2](https://github.com/VU-ASE/rover/commit/369c1d21e9c4e2335e4c4412a102ccffb643d681))

## [1.1.2](https://github.com/VU-ASE/rover/compare/v1.1.1...v1.1.2) (2025-05-21)


### Bug Fixes

* removed cargo modules from devcontainer ([f93a7be](https://github.com/VU-ASE/rover/commit/f93a7be6d2700ee86206234f9b4a77d26500e974))

## [1.1.1](https://github.com/VU-ASE/rover/compare/v1.1.0...v1.1.1) (2025-05-21)


### Bug Fixes

* added rover-local binary to CI ([c3c3ba3](https://github.com/VU-ASE/rover/commit/c3c3ba35e110476c0daa6549c49a36a0af7a3fa4))

## [1.1.0](https://github.com/VU-ASE/rover/compare/v1.0.2...v1.1.0) (2025-04-16)


### Features

* add debug support for generic int and generic float outputs ([5aedcee](https://github.com/VU-ASE/rover/commit/5aedceeac5edb627eb88800bef03e55bd383bb4a))
* add debugging support for energy sensor ([8885bdb](https://github.com/VU-ASE/rover/commit/8885bdba785e50a8bd01cba568563fddc96ef381))


### Bug Fixes

* broken link on dangers page ([be18f8a](https://github.com/VU-ASE/rover/commit/be18f8a3a3a22294b55df68761db780dae353495))

## [1.0.2](https://github.com/VU-ASE/rover/compare/v1.0.1...v1.0.2) (2025-04-01)


### Bug Fixes

* quick internet check before updating daemons ([e92f3ca](https://github.com/VU-ASE/rover/commit/e92f3ca5a63c4ef7c5071ccccd888486572de8ee))

## [1.0.1](https://github.com/VU-ASE/rover/compare/v1.0.1...v1.0.1) (2025-04-01)


### ⚠ BREAKING CHANGES

* release rover v1.0.0

### Features

* add alias "rover" to roverctl install script ([1ec8389](https://github.com/VU-ASE/rover/commit/1ec83893a0a7458aef2c8b41791c548ffda961df))
* add incompatibility warning to all endpoints that need it ([126849c](https://github.com/VU-ASE/rover/commit/126849cd89e743d089f2765f7261df24c30fd0eb))
* add installation location to service info view ([1a0ca7c](https://github.com/VU-ASE/rover/commit/1a0ca7c58c0a9b58ef3a0073babc92470ce99fe1))
* add passthrough back ([e07dc70](https://github.com/VU-ASE/rover/commit/e07dc7073c6b63dd9ed1c9894a11df2127e7d23c))
* add shortcuts item in navbar ([5aea4d3](https://github.com/VU-ASE/rover/commit/5aea4d3c7ec4458cd4aa6feb79e62398d5c1d860))
* added /emergency endpoint ([55df2dc](https://github.com/VU-ASE/rover/commit/55df2dcf7d51e36a6bec730cd36b2324de9cd45b))
* added Dockerfile for roverctl-web ([ec9733b](https://github.com/VU-ASE/rover/commit/ec9733b2f507d03d0ec405e6bcc2d82519406e6b))
* added manage script ([e1b1429](https://github.com/VU-ASE/rover/commit/e1b142989b1b0c0dfcaa55f0a496da56031288e4))
* added passthrough and debugging functionality ([90975ef](https://github.com/VU-ASE/rover/commit/90975ef34e1acf2ffc1f9ec740e18b600df7d561))
* api adaptation ([d36262f](https://github.com/VU-ASE/rover/commit/d36262f0e0d336e08cc1a3a6e2484ad4e52c2431))
* api clean up and addition of alias error (BBC) ([1d685a6](https://github.com/VU-ASE/rover/commit/1d685a6578ddc0b9c3b48fa9ffeb0abecfefad22))
* app config and overview page ([1c76d8c](https://github.com/VU-ASE/rover/commit/1c76d8cd42abe86a77e5133806ac97921a4da5d5))
* author command ([6ead41c](https://github.com/VU-ASE/rover/commit/6ead41cdbfb58ab0ce484cb627ae01bae0223da1))
* build command ([ea610eb](https://github.com/VU-ASE/rover/commit/ea610eb0171be57fbfd41ca2991bc6a1d00dedea))
* calibration command for actuator service ([a549e94](https://github.com/VU-ASE/rover/commit/a549e9499f9b629e89b40ea6db1fc4e1e6f025c4))
* CLI approach to roverctl ([bb78304](https://github.com/VU-ASE/rover/commit/bb78304cf2c6fcda4b2ce95bba51fcd183ee6ce3))
* dangers page ([a296ad3](https://github.com/VU-ASE/rover/commit/a296ad3ccb2d682d7a881c057d5a60aec36a9888))
* default ASE pipeline updates and installs ([0abcdb0](https://github.com/VU-ASE/rover/commit/0abcdb0021d0c2c4158a52e479684615864518af))
* dialog and upgrad/downgrade roverctl on version mismatch ([5d25e52](https://github.com/VU-ASE/rover/commit/5d25e52df3a75dfad549aca7e929f39d0755aa6f))
* emergency command ([2b883cd](https://github.com/VU-ASE/rover/commit/2b883cd7ded78acf211950161ec854884722badd))
* enable and install debug logic from roverctl-web ([a760e30](https://github.com/VU-ASE/rover/commit/a760e30d6991e0322399aeb4d65beca78a06dcc6))
* error types for the configuration endpoint ([016a94b](https://github.com/VU-ASE/rover/commit/016a94bd55a8643d300bf94d1c88c54bdb60512b))
* github actions workflows to build and push docker images ([5c4b035](https://github.com/VU-ASE/rover/commit/5c4b035329c9998331228113913a28fcc72b7552))
* group test scripts ([d073bfd](https://github.com/VU-ASE/rover/commit/d073bfdecdd198ae74f7b697144ed7dee806826c))
* helpful toasts and warning messages when pipeline starts/stops ([27074f6](https://github.com/VU-ASE/rover/commit/27074f67ca459497ca9105b6ec2c419f70a2d49b))
* implemented in place update API ([44a8c73](https://github.com/VU-ASE/rover/commit/44a8c73ffd2ad93d22056335e78b3a8ff4a5da39))
* improved error printing with multiline support ([8c64c2c](https://github.com/VU-ASE/rover/commit/8c64c2ce5443774ab1571658d5f57e932cb426b1))
* improved error reporting on upload failure ([b657f17](https://github.com/VU-ASE/rover/commit/b657f17d60bdaf464a2e998ddbb1541fead96141))
* info command ([478704e](https://github.com/VU-ASE/rover/commit/478704ec6fa6d5998f16da28bf9116df0c3dac45))
* info target ([0bc9cb0](https://github.com/VU-ASE/rover/commit/0bc9cb0ecfd0d6ffe2aac31c12717839d7ede73f))
* install command ([955fea9](https://github.com/VU-ASE/rover/commit/955fea970534b51d0e8e31abef62600f776475c1))
* install script that allows for installing a specific version ([f5bfcaf](https://github.com/VU-ASE/rover/commit/f5bfcaf2bada29e9e8732b61d79c93a53c3331d9))
* install script with symlink instead of alias ([22ccd67](https://github.com/VU-ASE/rover/commit/22ccd6786c112cbde2979064db8ae6ef89fd5c46))
* keyboard shortcuts for emergency, stopping pipelin and debug pause ([ad50f10](https://github.com/VU-ASE/rover/commit/ad50f1014705f1631c973928ff06922c74733514))
* link source code and docs for official services ([3ac9131](https://github.com/VU-ASE/rover/commit/3ac9131e4fce750e5475ec18167cf9feb83109c1))
* logs endpoint ([f1b3077](https://github.com/VU-ASE/rover/commit/f1b3077fb46bcc396dc9a7c1d32c96b8a66537d7))
* manage page (wip) and status overlay ([e9132ed](https://github.com/VU-ASE/rover/commit/e9132ed41535b49fbc10353c023d6935edf82640))
* merge passthrough into roverctl, debugging now works ([8f6c88f](https://github.com/VU-ASE/rover/commit/8f6c88f43323856971c7f6f129c47a0d23ba3a1e))
* pin version dependency between roverctl-roverd and warn if broken ([c8e98dc](https://github.com/VU-ASE/rover/commit/c8e98dcc4b89d079625197e91685f5b085be4bfd))
* pipeline enable/disable endpoints ([148fb7e](https://github.com/VU-ASE/rover/commit/148fb7eb16da42ae3af8ce568a767419260d9704))
* pipeline endpoint ([d0beca5](https://github.com/VU-ASE/rover/commit/d0beca50b2d01d0cedaada7833edd6a8bcba1053))
* release rover v1.0.0 ([292c229](https://github.com/VU-ASE/rover/commit/292c22939aeb4b8592c2dbe8ac1566c350d83b1a))
* reset calibration values on trimming ([ec8eb15](https://github.com/VU-ASE/rover/commit/ec8eb15b7efc09a2b245a1424f21fb4bb51beb9d))
* roverctl release ([143e8fa](https://github.com/VU-ASE/rover/commit/143e8faacb3775eee94b1835305c4a30def6d5a0))
* roverctl service init endpoint ([b8a9027](https://github.com/VU-ASE/rover/commit/b8a90274bc33f681fc5eaf97cb8adf04b828122f))
* roverctl-web startup through docker ([e3ca9a5](https://github.com/VU-ASE/rover/commit/e3ca9a54be65233a5f7a8f8fa8b915c9d042b724))
* service delete endpoint ([185362f](https://github.com/VU-ASE/rover/commit/185362f0f91d3af1cac2b0abe58e4741260740f9))
* service information ([c0119a5](https://github.com/VU-ASE/rover/commit/c0119a5c4715cb39d86fdecdf856ce3d48127595))
* service installation ([246a452](https://github.com/VU-ASE/rover/commit/246a452e0ae9b4d8489b1f852d1a425c7137e895))
* service selection toggles and ordening ([6ad02f4](https://github.com/VU-ASE/rover/commit/6ad02f49222b35caec422dc1e2c486276a50aa65))
* service_as feature ([782ab78](https://github.com/VU-ASE/rover/commit/782ab7874c1eb6eec0e58ac73283f2f9ab940a5b))
* services endpoint ([a5e11a4](https://github.com/VU-ASE/rover/commit/a5e11a46ca7d8c5a20345a356f32943e86ba50f0))
* shorthand aliasses for most commands ([3c270ce](https://github.com/VU-ASE/rover/commit/3c270ce7e0fbff41f1273a2cedd927f88d24fb99))
* show transceiver in roverctl-web ([be2cb92](https://github.com/VU-ASE/rover/commit/be2cb9274d7bed0c2d4e890b33018c8f82fbd652))
* show which service crashed the pipeline ([8591295](https://github.com/VU-ASE/rover/commit/85912954080cdb6654596702b195fb1c2b064ded))
* shutdown command ([4a124b7](https://github.com/VU-ASE/rover/commit/4a124b7ddec97a31e733b064745a4374e4cd71fc))
* start roverctl-web container ([9329015](https://github.com/VU-ASE/rover/commit/9329015b9cb86a8db86bc756052327fdf170785e))
* starting, stopping and modifying pipelines ([e54a176](https://github.com/VU-ASE/rover/commit/e54a1762fdc686ec0236874c1169ea9c20014bff))
* unified roverd and roverctl under 'rover' repository ([f893365](https://github.com/VU-ASE/rover/commit/f893365bda94929d5cfa78a4a5d4407ba4d87d95))
* unified toasts ([ded557b](https://github.com/VU-ASE/rover/commit/ded557be905b07b688a2fe9f6098a028543ec8bb))
* update endpoints ([e1422b4](https://github.com/VU-ASE/rover/commit/e1422b40ed2e25e98eb094bc088154d896a28caf))
* update target ([0f3cc2a](https://github.com/VU-ASE/rover/commit/0f3cc2a340b5a78a3c04fc363941608fa9622613))
* upgrade roverctl to go 1.22.12 ([6a19635](https://github.com/VU-ASE/rover/commit/6a1963502affb37a09328151d1c55622ffb14d0f))
* use mDNS to resolve hostnames, allow proxy IP override ([2cb436e](https://github.com/VU-ASE/rover/commit/2cb436e2038e2ef1e5d722913f0d794ce9d60bfe))
* warn users if roverctl mismatches roverd and install appr roverctl ([868e3c1](https://github.com/VU-ASE/rover/commit/868e3c1db16c7dc0955121ae022ded73ff152f0a))


### Bug Fixes

* added service-template for cpp ([cf5fd0c](https://github.com/VU-ASE/rover/commit/cf5fd0c04b646dfef25138bbbb54220c03a2722e))
* added smoother steering to testing scripts ([681d881](https://github.com/VU-ASE/rover/commit/681d881c2147039cd8d2954098833fd0110bc080))
* auto-update roverd in test ([9b53863](https://github.com/VU-ASE/rover/commit/9b53863f55a2baf712b25f46099b1d2bbd7b7ff3))
* broken links ([7a4d2aa](https://github.com/VU-ASE/rover/commit/7a4d2aa8ee9bc87d5f06ea3c1f55a9f324cc57a7))
* build to ./bin ([dba0f2b](https://github.com/VU-ASE/rover/commit/dba0f2b1cfd4deaa257beb7c0bbe241b02879f5d))
* ci testing ([269b7a3](https://github.com/VU-ASE/rover/commit/269b7a3b9ece651b79b22145db9eccb1cf018fa9))
* consistent style for services ([a9f5fbe](https://github.com/VU-ASE/rover/commit/a9f5fbef311b50ce30ca3891a9c5591ec3c12d91))
* correct link if upload fails ([d229edb](https://github.com/VU-ASE/rover/commit/d229edbeca49c5684dd881e97dc9d02bf9cbb9fc))
* correct timing of service error messages to actual moment of crash ([68f158f](https://github.com/VU-ASE/rover/commit/68f158f9404125399fc18499d855f04f6ab054f0))
* create openapi directory if it does not exist ([1f881b8](https://github.com/VU-ASE/rover/commit/1f881b8846f38ceea597ab31be031ae7c60ecf96))
* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* disable transceiver service when debug mode is not enabled ([46f41e5](https://github.com/VU-ASE/rover/commit/46f41e5f75025b3dd3515852eb42f69f88afccea))
* display warning if author is not set in info command ([b2c5eaa](https://github.com/VU-ASE/rover/commit/b2c5eaa5e866c87066cdd2eddc1aecd6f680d493))
* do not build roverd on self-hosted ([af62651](https://github.com/VU-ASE/rover/commit/af626516335e3f1c2a5682aaa8fb54b5f72deb5c))
* do not error on existing connection being reused ([6e357bf](https://github.com/VU-ASE/rover/commit/6e357bfd6dc7c008d09f0717992a3619f7f2240c))
* do not publish openapi files ([f7f7ff0](https://github.com/VU-ASE/rover/commit/f7f7ff039253c5ebfc893fa0d53b5daad304968d))
* do not reset calibration value for normal pipeline testing ([be21d46](https://github.com/VU-ASE/rover/commit/be21d462d992f50a90e93b473417e6a0f026b3f3))
* do not show "shortcuts" tab as active on debugpage ([822b084](https://github.com/VU-ASE/rover/commit/822b084f1a15a00c34a12a162d8c618e61bc5c65))
* do NOT update roverd in the test script ([b5f1cb0](https://github.com/VU-ASE/rover/commit/b5f1cb0ea34f527117f6aee6a8d071002786c52d))
* docs and install script for roverctl ([a069205](https://github.com/VU-ASE/rover/commit/a069205b726187a5ad7ad99c8f00a75f7569ce9f))
* emergency stop ([2b7bdf8](https://github.com/VU-ASE/rover/commit/2b7bdf85012d4b5e736a5bf3543a7f5856bc9855))
* error on roverctl install.sh if release does not exist or no binary ([1880064](https://github.com/VU-ASE/rover/commit/1880064e54514a502ae07bd21f690b8d953a7eea))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* extract artifact downloads first ([37762b7](https://github.com/VU-ASE/rover/commit/37762b77e75becc3a8fb960e83ddd72f47305729))
* extract from correct location ([d4106ba](https://github.com/VU-ASE/rover/commit/d4106ba17d7e8ec4d7fea3f3827aded9d4a25d0a))
* gitignore was missing readme ([44c64ff](https://github.com/VU-ASE/rover/commit/44c64ffed3c5d317704599bfb95af5ace86af280))
* hard-coded version ([058573d](https://github.com/VU-ASE/rover/commit/058573d9cf67a814ef554712b4ad610eaabcb0b4))
* hide updater screen when all services are installed ([f4dce16](https://github.com/VU-ASE/rover/commit/f4dce16e2ae91666336fa361d6c2d15db813c380))
* image push only succeeds on --verbose flag, port checking ([b0d5998](https://github.com/VU-ASE/rover/commit/b0d599895c6cc1f3365eca1aee7ceadbadb6e13e))
* improved upload description and example usage ([248002c](https://github.com/VU-ASE/rover/commit/248002c3d1c96494a91577621a270ea0a3a68cea))
* incorrect number of arguments for service init script ([49981a0](https://github.com/VU-ASE/rover/commit/49981a094084063c1f584caf0f3df2a25e13aeca))
* infinite redirect ([0f63b01](https://github.com/VU-ASE/rover/commit/0f63b019402c498c496f4cec1bbf9f483de3065a))
* load list of services from roverd in debug window ([a8099bb](https://github.com/VU-ASE/rover/commit/a8099bb929543b3dfde14b06bc61115c880f3ceb))
* make roverctl compatible with new API spec ([6eb675c](https://github.com/VU-ASE/rover/commit/6eb675c87be7b49d31ca3405906767df359405f3))
* merge service name and service "as" name in tuning overview ([3675c0e](https://github.com/VU-ASE/rover/commit/3675c0e5159f0e503c8ff65a596c5f757263a9c5))
* more helpful messages to enable debug mode ([106e3ea](https://github.com/VU-ASE/rover/commit/106e3ea929bc7943948e77f07074171e4a31ad55))
* more lenient timeouts for roverd downloads ([a89b302](https://github.com/VU-ASE/rover/commit/a89b302fd97b48ce176d1f27becee737a67ac803))
* move github.com/VU-ASE/roverctl to github.com/VU-ASE/rover/roverctl ([a35e0dc](https://github.com/VU-ASE/rover/commit/a35e0dc12895750545256432284054e560e9eb65))
* navbar type error ([d66c700](https://github.com/VU-ASE/rover/commit/d66c70009adc0a222eba5967af53d168de2f44af))
* only show services when the pipeline is started (in debug mode) ([e22facf](https://github.com/VU-ASE/rover/commit/e22facf86deb8c7b305238dcc704b6aab7fc741b))
* open-api-gen debugging ([013bce9](https://github.com/VU-ASE/rover/commit/013bce9f4a0838ccf91ddb53e429d27cb7132178))
* permission issues for building ([8f82d02](https://github.com/VU-ASE/rover/commit/8f82d02b4b4bcd6a5930c5faa8fd504042a69f0b))
* permissions on github actions ([0b7225a](https://github.com/VU-ASE/rover/commit/0b7225ae596dd3d1e1a29a46ad26b2f6e256d9dd))
* pipeline not loading after viewing debug screen ([67787c4](https://github.com/VU-ASE/rover/commit/67787c480eb8d8c92381433008e21b2297b5b6eb))
* point install scripts to rover repository (was roverctl) ([b139324](https://github.com/VU-ASE/rover/commit/b139324d303d27c88f2f8d91d0aaf1b78a5f714a))
* poll pipeline status ([a02bbf6](https://github.com/VU-ASE/rover/commit/a02bbf698dbd8888b4f8efcf82dfb2fdb7eca9a8))
* proxy server deadlock on repeated transceiver crashes ([577a1a2](https://github.com/VU-ASE/rover/commit/577a1a29775debc7e75526abbc9f799e9b3db0ec))
* public link for schematic in docs ([b26decf](https://github.com/VU-ASE/rover/commit/b26decf57eb08b15e0fd4c2c04e81d1f9b3a1fdd))
* release assets on all success ([4448019](https://github.com/VU-ASE/rover/commit/44480199fcac3cc9c81e317ef7b34c270641ec6a))
* remove unused github action jobs ([04ef005](https://github.com/VU-ASE/rover/commit/04ef0051fa635a52e6fbc521c6ca915b58c1423e))
* remove unused view ([31f91cd](https://github.com/VU-ASE/rover/commit/31f91cd1bd482a2a1bb05af107ebb19b153d1de5))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* rename passthrough to appropriate folder ([5c1d41e](https://github.com/VU-ASE/rover/commit/5c1d41e8f3a1bb3be425841ede4c9a80fe2c2861))
* retain debug toggle information and don't show unusable tuning ([9d78947](https://github.com/VU-ASE/rover/commit/9d789474196bc613cea17b1da0a775e8669e9d12))
* returning service_as optimistically for all endpoints ([a019fe1](https://github.com/VU-ASE/rover/commit/a019fe1bee5a0a5e591911b389ee599b647e0440))
* roverctl --force warning message ([5384c70](https://github.com/VU-ASE/rover/commit/5384c704cac253524bc7d433d4b260602e1c3632))
* roverctl install.sh, verify downloaded binary is valid ([381a5ed](https://github.com/VU-ASE/rover/commit/381a5ed2e809710a882a65e4affd3fb658ecb2e2))
* roverctl releae ([2d39d11](https://github.com/VU-ASE/rover/commit/2d39d11c5d2d80bce19a9002462acdf65f953018))
* roverctl release 1.0.0 ([ec304d4](https://github.com/VU-ASE/rover/commit/ec304d46cf0205bed12551c316e67ad69d277c0b))
* roverctl replace FQN types ([f6a9c44](https://github.com/VU-ASE/rover/commit/f6a9c44dee576b4ba63f3e44ca5a07704ead245b))
* roverctl-roverd update API endpoint not working ([9fd776f](https://github.com/VU-ASE/rover/commit/9fd776f4616d2c66e59a66cb0c2b554620a72522))
* roverctl, build openapi files in Dockerfile ([af01eb6](https://github.com/VU-ASE/rover/commit/af01eb620b565518860cb84d5665595788b0de71))
* roverd early check for root ([451f7eb](https://github.com/VU-ASE/rover/commit/451f7eb4ab031fe82dc0d7b19e2a7377122a2e5f))
* roverd log lines ([97adc00](https://github.com/VU-ASE/rover/commit/97adc0044f551ce2aa9ea1e7db9beeae24e43ab7))
* roverd not propagating stream outputs to transceiver if no consumer ([0818f38](https://github.com/VU-ASE/rover/commit/0818f38c0b4353bb5078a3c8b904be8a9e5f3c38))
* roverd shows all exit codes now ([2de0226](https://github.com/VU-ASE/rover/commit/2de022691173795eac6a63eec9787d554f466958))
* roverd shutsdown pipeline on termination ([83ea1dc](https://github.com/VU-ASE/rover/commit/83ea1dce347970d9a265472b62a18031177b46eb))
* roverd unclosed delimiter ([17d8e95](https://github.com/VU-ASE/rover/commit/17d8e95a9b322a205ffe46c5dd367459f59c9aef))
* run open-api generator as devuser ([e8e84c8](https://github.com/VU-ASE/rover/commit/e8e84c85578d578299798f469782f04e9acff0f0))
* service logs cut-off point makes them unreadable ([412a8ee](https://github.com/VU-ASE/rover/commit/412a8ee5eb7520daee8fbef481b0271eadab9470))
* service yamls can now specify a bash command for run ([6022928](https://github.com/VU-ASE/rover/commit/60229285d294b95698faef9958704e3d8c81ff5e))
* show correct fix for updating roverd on mismatch ([a49144f](https://github.com/VU-ASE/rover/commit/a49144feec173a050d84a6bfee482bcb7fe7afc3))
* show correct panels for selected services ([5ba569e](https://github.com/VU-ASE/rover/commit/5ba569e06f205686fd5baa2dd7ff0ae765d06bb7))
* show debug toggle correctly ([3ea63f8](https://github.com/VU-ASE/rover/commit/3ea63f8a79ee4fe7ca0836407fb96d6f501fd94d))
* show logs of last run when selected ([0cca2da](https://github.com/VU-ASE/rover/commit/0cca2dab6e762f6fb8916ed3cce3175f17262ec0))
* show pipeline edges accurately ([ff53e34](https://github.com/VU-ASE/rover/commit/ff53e3424e6c8644bd943c3d181c080cf69f6a74))
* show service crash exit code on transition from started to stopped ([e43cb26](https://github.com/VU-ASE/rover/commit/e43cb260ef1df98869502afda7731fe188257beb))
* sort services alphabetically. Always show vu-ase first ([fb6a58b](https://github.com/VU-ASE/rover/commit/fb6a58b9419af46bb1b06792e548bdc9431fb6da))
* specifying value types in service.yaml is required ([e8757a3](https://github.com/VU-ASE/rover/commit/e8757a3edeb67edcccb0cd5d6909d15879d6f886))
* testing ci again ([573bb0f](https://github.com/VU-ASE/rover/commit/573bb0f5aafb6ca7b18208a4c71b1a14724de484))
* transceiver output address should use `*` instead of `localhost` ([19264da](https://github.com/VU-ASE/rover/commit/19264dac0d2c18e4ff8485889e4f9dd0dd0fddab))
* update checking for default pipeline ([3f67220](https://github.com/VU-ASE/rover/commit/3f672204d7faaeae6ec4a3d2ec141ab52af4dc55))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))
* use correct version of action/artifact ([300beda](https://github.com/VU-ASE/rover/commit/300beda2031241877d4101279e5f9ec5cf8dcd07))
* use docker buildx for builds ([aa4893b](https://github.com/VU-ASE/rover/commit/aa4893b9e97492b82198164711f7cca2e534f094))
* use dynamic env variables for configuration ([25d5c88](https://github.com/VU-ASE/rover/commit/25d5c88b217c9040aea12686a4b08b05443a8e03))
* use lowercase for GHCR tag ([e62a583](https://github.com/VU-ASE/rover/commit/e62a58302e17c7ca9739349b60fc654cce7f7da3))
* use mdns for tests ([bb59b70](https://github.com/VU-ASE/rover/commit/bb59b70cbbbc7e8a31fdf4e93a85f63346c22e5f))
* using mv instead of os.Rename ([696ecc4](https://github.com/VU-ASE/rover/commit/696ecc440505fb39c3ec5cf814b55786b9fb5bf5))
* version mismatch checking ([4ba63b9](https://github.com/VU-ASE/rover/commit/4ba63b901f88055e68db2a7234ca17659588798c))
* version update script now accepts "v" prefix ([362c24c](https://github.com/VU-ASE/rover/commit/362c24c0b3264ce7d3aef7fdfabbfb6bea72d45a))
* version utility and correct comparison ([9251512](https://github.com/VU-ASE/rover/commit/92515124a450c2edd2ad8aed10a20b7b42ba4b9c))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))
* release 0.0.1 ([abd806e](https://github.com/VU-ASE/rover/commit/abd806e59353ee5ea9da82a80612ff9ece653285))
* release 0.0.5 ([aa2ed11](https://github.com/VU-ASE/rover/commit/aa2ed11e18453f8a6edc3f82cd503e38288e7060))
* release 1.0.0 ([b1c1497](https://github.com/VU-ASE/rover/commit/b1c14978496488dac92350a6c75cb27156cd017a))
* release 1.0.0 ([55d09fd](https://github.com/VU-ASE/rover/commit/55d09fd29d0bc18c528ce63af45a2c10c3428f8d))
* release 1.0.1 ([8b1afab](https://github.com/VU-ASE/rover/commit/8b1afab7c598761bd0943599ef8b83a50ba24e49))
* release 1.0.1 ([ea40d3b](https://github.com/VU-ASE/rover/commit/ea40d3bcfe9ac8d3147e4b2873d05a1b305d348f))

## [1.0.1](https://github.com/VU-ASE/rover/compare/v1.0.0...v1.0.1) (2025-04-01)


### ⚠ BREAKING CHANGES

* release rover v1.0.0

### Features

* add alias "rover" to roverctl install script ([1ec8389](https://github.com/VU-ASE/rover/commit/1ec83893a0a7458aef2c8b41791c548ffda961df))
* add incompatibility warning to all endpoints that need it ([126849c](https://github.com/VU-ASE/rover/commit/126849cd89e743d089f2765f7261df24c30fd0eb))
* add installation location to service info view ([1a0ca7c](https://github.com/VU-ASE/rover/commit/1a0ca7c58c0a9b58ef3a0073babc92470ce99fe1))
* add passthrough back ([e07dc70](https://github.com/VU-ASE/rover/commit/e07dc7073c6b63dd9ed1c9894a11df2127e7d23c))
* add shortcuts item in navbar ([5aea4d3](https://github.com/VU-ASE/rover/commit/5aea4d3c7ec4458cd4aa6feb79e62398d5c1d860))
* added /emergency endpoint ([55df2dc](https://github.com/VU-ASE/rover/commit/55df2dcf7d51e36a6bec730cd36b2324de9cd45b))
* added Dockerfile for roverctl-web ([ec9733b](https://github.com/VU-ASE/rover/commit/ec9733b2f507d03d0ec405e6bcc2d82519406e6b))
* added manage script ([e1b1429](https://github.com/VU-ASE/rover/commit/e1b142989b1b0c0dfcaa55f0a496da56031288e4))
* added passthrough and debugging functionality ([90975ef](https://github.com/VU-ASE/rover/commit/90975ef34e1acf2ffc1f9ec740e18b600df7d561))
* api adaptation ([d36262f](https://github.com/VU-ASE/rover/commit/d36262f0e0d336e08cc1a3a6e2484ad4e52c2431))
* api clean up and addition of alias error (BBC) ([1d685a6](https://github.com/VU-ASE/rover/commit/1d685a6578ddc0b9c3b48fa9ffeb0abecfefad22))
* app config and overview page ([1c76d8c](https://github.com/VU-ASE/rover/commit/1c76d8cd42abe86a77e5133806ac97921a4da5d5))
* author command ([6ead41c](https://github.com/VU-ASE/rover/commit/6ead41cdbfb58ab0ce484cb627ae01bae0223da1))
* build command ([ea610eb](https://github.com/VU-ASE/rover/commit/ea610eb0171be57fbfd41ca2991bc6a1d00dedea))
* calibration command for actuator service ([a549e94](https://github.com/VU-ASE/rover/commit/a549e9499f9b629e89b40ea6db1fc4e1e6f025c4))
* CLI approach to roverctl ([bb78304](https://github.com/VU-ASE/rover/commit/bb78304cf2c6fcda4b2ce95bba51fcd183ee6ce3))
* dangers page ([a296ad3](https://github.com/VU-ASE/rover/commit/a296ad3ccb2d682d7a881c057d5a60aec36a9888))
* default ASE pipeline updates and installs ([0abcdb0](https://github.com/VU-ASE/rover/commit/0abcdb0021d0c2c4158a52e479684615864518af))
* dialog and upgrad/downgrade roverctl on version mismatch ([5d25e52](https://github.com/VU-ASE/rover/commit/5d25e52df3a75dfad549aca7e929f39d0755aa6f))
* emergency command ([2b883cd](https://github.com/VU-ASE/rover/commit/2b883cd7ded78acf211950161ec854884722badd))
* enable and install debug logic from roverctl-web ([a760e30](https://github.com/VU-ASE/rover/commit/a760e30d6991e0322399aeb4d65beca78a06dcc6))
* error types for the configuration endpoint ([016a94b](https://github.com/VU-ASE/rover/commit/016a94bd55a8643d300bf94d1c88c54bdb60512b))
* github actions workflows to build and push docker images ([5c4b035](https://github.com/VU-ASE/rover/commit/5c4b035329c9998331228113913a28fcc72b7552))
* group test scripts ([d073bfd](https://github.com/VU-ASE/rover/commit/d073bfdecdd198ae74f7b697144ed7dee806826c))
* helpful toasts and warning messages when pipeline starts/stops ([27074f6](https://github.com/VU-ASE/rover/commit/27074f67ca459497ca9105b6ec2c419f70a2d49b))
* implemented in place update API ([44a8c73](https://github.com/VU-ASE/rover/commit/44a8c73ffd2ad93d22056335e78b3a8ff4a5da39))
* improved error printing with multiline support ([8c64c2c](https://github.com/VU-ASE/rover/commit/8c64c2ce5443774ab1571658d5f57e932cb426b1))
* improved error reporting on upload failure ([b657f17](https://github.com/VU-ASE/rover/commit/b657f17d60bdaf464a2e998ddbb1541fead96141))
* info command ([478704e](https://github.com/VU-ASE/rover/commit/478704ec6fa6d5998f16da28bf9116df0c3dac45))
* info target ([0bc9cb0](https://github.com/VU-ASE/rover/commit/0bc9cb0ecfd0d6ffe2aac31c12717839d7ede73f))
* install command ([955fea9](https://github.com/VU-ASE/rover/commit/955fea970534b51d0e8e31abef62600f776475c1))
* install script that allows for installing a specific version ([f5bfcaf](https://github.com/VU-ASE/rover/commit/f5bfcaf2bada29e9e8732b61d79c93a53c3331d9))
* install script with symlink instead of alias ([22ccd67](https://github.com/VU-ASE/rover/commit/22ccd6786c112cbde2979064db8ae6ef89fd5c46))
* keyboard shortcuts for emergency, stopping pipelin and debug pause ([ad50f10](https://github.com/VU-ASE/rover/commit/ad50f1014705f1631c973928ff06922c74733514))
* link source code and docs for official services ([3ac9131](https://github.com/VU-ASE/rover/commit/3ac9131e4fce750e5475ec18167cf9feb83109c1))
* logs endpoint ([f1b3077](https://github.com/VU-ASE/rover/commit/f1b3077fb46bcc396dc9a7c1d32c96b8a66537d7))
* manage page (wip) and status overlay ([e9132ed](https://github.com/VU-ASE/rover/commit/e9132ed41535b49fbc10353c023d6935edf82640))
* merge passthrough into roverctl, debugging now works ([8f6c88f](https://github.com/VU-ASE/rover/commit/8f6c88f43323856971c7f6f129c47a0d23ba3a1e))
* pin version dependency between roverctl-roverd and warn if broken ([c8e98dc](https://github.com/VU-ASE/rover/commit/c8e98dcc4b89d079625197e91685f5b085be4bfd))
* pipeline enable/disable endpoints ([148fb7e](https://github.com/VU-ASE/rover/commit/148fb7eb16da42ae3af8ce568a767419260d9704))
* pipeline endpoint ([d0beca5](https://github.com/VU-ASE/rover/commit/d0beca50b2d01d0cedaada7833edd6a8bcba1053))
* release rover v1.0.0 ([292c229](https://github.com/VU-ASE/rover/commit/292c22939aeb4b8592c2dbe8ac1566c350d83b1a))
* reset calibration values on trimming ([ec8eb15](https://github.com/VU-ASE/rover/commit/ec8eb15b7efc09a2b245a1424f21fb4bb51beb9d))
* roverctl release ([143e8fa](https://github.com/VU-ASE/rover/commit/143e8faacb3775eee94b1835305c4a30def6d5a0))
* roverctl service init endpoint ([b8a9027](https://github.com/VU-ASE/rover/commit/b8a90274bc33f681fc5eaf97cb8adf04b828122f))
* roverctl-web startup through docker ([e3ca9a5](https://github.com/VU-ASE/rover/commit/e3ca9a54be65233a5f7a8f8fa8b915c9d042b724))
* service delete endpoint ([185362f](https://github.com/VU-ASE/rover/commit/185362f0f91d3af1cac2b0abe58e4741260740f9))
* service information ([c0119a5](https://github.com/VU-ASE/rover/commit/c0119a5c4715cb39d86fdecdf856ce3d48127595))
* service installation ([246a452](https://github.com/VU-ASE/rover/commit/246a452e0ae9b4d8489b1f852d1a425c7137e895))
* service selection toggles and ordening ([6ad02f4](https://github.com/VU-ASE/rover/commit/6ad02f49222b35caec422dc1e2c486276a50aa65))
* service_as feature ([782ab78](https://github.com/VU-ASE/rover/commit/782ab7874c1eb6eec0e58ac73283f2f9ab940a5b))
* services endpoint ([a5e11a4](https://github.com/VU-ASE/rover/commit/a5e11a46ca7d8c5a20345a356f32943e86ba50f0))
* shorthand aliasses for most commands ([3c270ce](https://github.com/VU-ASE/rover/commit/3c270ce7e0fbff41f1273a2cedd927f88d24fb99))
* show transceiver in roverctl-web ([be2cb92](https://github.com/VU-ASE/rover/commit/be2cb9274d7bed0c2d4e890b33018c8f82fbd652))
* show which service crashed the pipeline ([8591295](https://github.com/VU-ASE/rover/commit/85912954080cdb6654596702b195fb1c2b064ded))
* shutdown command ([4a124b7](https://github.com/VU-ASE/rover/commit/4a124b7ddec97a31e733b064745a4374e4cd71fc))
* start roverctl-web container ([9329015](https://github.com/VU-ASE/rover/commit/9329015b9cb86a8db86bc756052327fdf170785e))
* starting, stopping and modifying pipelines ([e54a176](https://github.com/VU-ASE/rover/commit/e54a1762fdc686ec0236874c1169ea9c20014bff))
* unified toasts ([ded557b](https://github.com/VU-ASE/rover/commit/ded557be905b07b688a2fe9f6098a028543ec8bb))
* update endpoints ([e1422b4](https://github.com/VU-ASE/rover/commit/e1422b40ed2e25e98eb094bc088154d896a28caf))
* update target ([0f3cc2a](https://github.com/VU-ASE/rover/commit/0f3cc2a340b5a78a3c04fc363941608fa9622613))
* upgrade roverctl to go 1.22.12 ([6a19635](https://github.com/VU-ASE/rover/commit/6a1963502affb37a09328151d1c55622ffb14d0f))
* use mDNS to resolve hostnames, allow proxy IP override ([2cb436e](https://github.com/VU-ASE/rover/commit/2cb436e2038e2ef1e5d722913f0d794ce9d60bfe))
* warn users if roverctl mismatches roverd and install appr roverctl ([868e3c1](https://github.com/VU-ASE/rover/commit/868e3c1db16c7dc0955121ae022ded73ff152f0a))


### Bug Fixes

* added service-template for cpp ([cf5fd0c](https://github.com/VU-ASE/rover/commit/cf5fd0c04b646dfef25138bbbb54220c03a2722e))
* added smoother steering to testing scripts ([681d881](https://github.com/VU-ASE/rover/commit/681d881c2147039cd8d2954098833fd0110bc080))
* auto-update roverd in test ([9b53863](https://github.com/VU-ASE/rover/commit/9b53863f55a2baf712b25f46099b1d2bbd7b7ff3))
* broken links ([7a4d2aa](https://github.com/VU-ASE/rover/commit/7a4d2aa8ee9bc87d5f06ea3c1f55a9f324cc57a7))
* build to ./bin ([dba0f2b](https://github.com/VU-ASE/rover/commit/dba0f2b1cfd4deaa257beb7c0bbe241b02879f5d))
* ci testing ([269b7a3](https://github.com/VU-ASE/rover/commit/269b7a3b9ece651b79b22145db9eccb1cf018fa9))
* consistent style for services ([a9f5fbe](https://github.com/VU-ASE/rover/commit/a9f5fbef311b50ce30ca3891a9c5591ec3c12d91))
* correct link if upload fails ([d229edb](https://github.com/VU-ASE/rover/commit/d229edbeca49c5684dd881e97dc9d02bf9cbb9fc))
* correct timing of service error messages to actual moment of crash ([68f158f](https://github.com/VU-ASE/rover/commit/68f158f9404125399fc18499d855f04f6ab054f0))
* create openapi directory if it does not exist ([1f881b8](https://github.com/VU-ASE/rover/commit/1f881b8846f38ceea597ab31be031ae7c60ecf96))
* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* disable transceiver service when debug mode is not enabled ([46f41e5](https://github.com/VU-ASE/rover/commit/46f41e5f75025b3dd3515852eb42f69f88afccea))
* display warning if author is not set in info command ([b2c5eaa](https://github.com/VU-ASE/rover/commit/b2c5eaa5e866c87066cdd2eddc1aecd6f680d493))
* do not build roverd on self-hosted ([af62651](https://github.com/VU-ASE/rover/commit/af626516335e3f1c2a5682aaa8fb54b5f72deb5c))
* do not error on existing connection being reused ([6e357bf](https://github.com/VU-ASE/rover/commit/6e357bfd6dc7c008d09f0717992a3619f7f2240c))
* do not publish openapi files ([f7f7ff0](https://github.com/VU-ASE/rover/commit/f7f7ff039253c5ebfc893fa0d53b5daad304968d))
* do not reset calibration value for normal pipeline testing ([be21d46](https://github.com/VU-ASE/rover/commit/be21d462d992f50a90e93b473417e6a0f026b3f3))
* do not show "shortcuts" tab as active on debugpage ([822b084](https://github.com/VU-ASE/rover/commit/822b084f1a15a00c34a12a162d8c618e61bc5c65))
* do NOT update roverd in the test script ([b5f1cb0](https://github.com/VU-ASE/rover/commit/b5f1cb0ea34f527117f6aee6a8d071002786c52d))
* docs and install script for roverctl ([a069205](https://github.com/VU-ASE/rover/commit/a069205b726187a5ad7ad99c8f00a75f7569ce9f))
* emergency stop ([2b7bdf8](https://github.com/VU-ASE/rover/commit/2b7bdf85012d4b5e736a5bf3543a7f5856bc9855))
* error on roverctl install.sh if release does not exist or no binary ([1880064](https://github.com/VU-ASE/rover/commit/1880064e54514a502ae07bd21f690b8d953a7eea))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* extract artifact downloads first ([37762b7](https://github.com/VU-ASE/rover/commit/37762b77e75becc3a8fb960e83ddd72f47305729))
* extract from correct location ([d4106ba](https://github.com/VU-ASE/rover/commit/d4106ba17d7e8ec4d7fea3f3827aded9d4a25d0a))
* gitignore was missing readme ([44c64ff](https://github.com/VU-ASE/rover/commit/44c64ffed3c5d317704599bfb95af5ace86af280))
* hard-coded version ([058573d](https://github.com/VU-ASE/rover/commit/058573d9cf67a814ef554712b4ad610eaabcb0b4))
* hide updater screen when all services are installed ([f4dce16](https://github.com/VU-ASE/rover/commit/f4dce16e2ae91666336fa361d6c2d15db813c380))
* image push only succeeds on --verbose flag, port checking ([b0d5998](https://github.com/VU-ASE/rover/commit/b0d599895c6cc1f3365eca1aee7ceadbadb6e13e))
* improved upload description and example usage ([248002c](https://github.com/VU-ASE/rover/commit/248002c3d1c96494a91577621a270ea0a3a68cea))
* incorrect number of arguments for service init script ([49981a0](https://github.com/VU-ASE/rover/commit/49981a094084063c1f584caf0f3df2a25e13aeca))
* infinite redirect ([0f63b01](https://github.com/VU-ASE/rover/commit/0f63b019402c498c496f4cec1bbf9f483de3065a))
* load list of services from roverd in debug window ([a8099bb](https://github.com/VU-ASE/rover/commit/a8099bb929543b3dfde14b06bc61115c880f3ceb))
* make roverctl compatible with new API spec ([6eb675c](https://github.com/VU-ASE/rover/commit/6eb675c87be7b49d31ca3405906767df359405f3))
* merge service name and service "as" name in tuning overview ([3675c0e](https://github.com/VU-ASE/rover/commit/3675c0e5159f0e503c8ff65a596c5f757263a9c5))
* more helpful messages to enable debug mode ([106e3ea](https://github.com/VU-ASE/rover/commit/106e3ea929bc7943948e77f07074171e4a31ad55))
* more lenient timeouts for roverd downloads ([a89b302](https://github.com/VU-ASE/rover/commit/a89b302fd97b48ce176d1f27becee737a67ac803))
* move github.com/VU-ASE/roverctl to github.com/VU-ASE/rover/roverctl ([a35e0dc](https://github.com/VU-ASE/rover/commit/a35e0dc12895750545256432284054e560e9eb65))
* navbar type error ([d66c700](https://github.com/VU-ASE/rover/commit/d66c70009adc0a222eba5967af53d168de2f44af))
* only show services when the pipeline is started (in debug mode) ([e22facf](https://github.com/VU-ASE/rover/commit/e22facf86deb8c7b305238dcc704b6aab7fc741b))
* open-api-gen debugging ([013bce9](https://github.com/VU-ASE/rover/commit/013bce9f4a0838ccf91ddb53e429d27cb7132178))
* permission issues for building ([8f82d02](https://github.com/VU-ASE/rover/commit/8f82d02b4b4bcd6a5930c5faa8fd504042a69f0b))
* permissions on github actions ([0b7225a](https://github.com/VU-ASE/rover/commit/0b7225ae596dd3d1e1a29a46ad26b2f6e256d9dd))
* pipeline not loading after viewing debug screen ([67787c4](https://github.com/VU-ASE/rover/commit/67787c480eb8d8c92381433008e21b2297b5b6eb))
* point install scripts to rover repository (was roverctl) ([b139324](https://github.com/VU-ASE/rover/commit/b139324d303d27c88f2f8d91d0aaf1b78a5f714a))
* poll pipeline status ([a02bbf6](https://github.com/VU-ASE/rover/commit/a02bbf698dbd8888b4f8efcf82dfb2fdb7eca9a8))
* proxy server deadlock on repeated transceiver crashes ([577a1a2](https://github.com/VU-ASE/rover/commit/577a1a29775debc7e75526abbc9f799e9b3db0ec))
* public link for schematic in docs ([b26decf](https://github.com/VU-ASE/rover/commit/b26decf57eb08b15e0fd4c2c04e81d1f9b3a1fdd))
* release assets on all success ([4448019](https://github.com/VU-ASE/rover/commit/44480199fcac3cc9c81e317ef7b34c270641ec6a))
* remove unused github action jobs ([04ef005](https://github.com/VU-ASE/rover/commit/04ef0051fa635a52e6fbc521c6ca915b58c1423e))
* remove unused view ([31f91cd](https://github.com/VU-ASE/rover/commit/31f91cd1bd482a2a1bb05af107ebb19b153d1de5))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* rename passthrough to appropriate folder ([5c1d41e](https://github.com/VU-ASE/rover/commit/5c1d41e8f3a1bb3be425841ede4c9a80fe2c2861))
* retain debug toggle information and don't show unusable tuning ([9d78947](https://github.com/VU-ASE/rover/commit/9d789474196bc613cea17b1da0a775e8669e9d12))
* returning service_as optimistically for all endpoints ([a019fe1](https://github.com/VU-ASE/rover/commit/a019fe1bee5a0a5e591911b389ee599b647e0440))
* roverctl --force warning message ([5384c70](https://github.com/VU-ASE/rover/commit/5384c704cac253524bc7d433d4b260602e1c3632))
* roverctl install.sh, verify downloaded binary is valid ([381a5ed](https://github.com/VU-ASE/rover/commit/381a5ed2e809710a882a65e4affd3fb658ecb2e2))
* roverctl releae ([2d39d11](https://github.com/VU-ASE/rover/commit/2d39d11c5d2d80bce19a9002462acdf65f953018))
* roverctl release 1.0.0 ([ec304d4](https://github.com/VU-ASE/rover/commit/ec304d46cf0205bed12551c316e67ad69d277c0b))
* roverctl replace FQN types ([f6a9c44](https://github.com/VU-ASE/rover/commit/f6a9c44dee576b4ba63f3e44ca5a07704ead245b))
* roverctl-roverd update API endpoint not working ([9fd776f](https://github.com/VU-ASE/rover/commit/9fd776f4616d2c66e59a66cb0c2b554620a72522))
* roverctl, build openapi files in Dockerfile ([af01eb6](https://github.com/VU-ASE/rover/commit/af01eb620b565518860cb84d5665595788b0de71))
* roverd early check for root ([451f7eb](https://github.com/VU-ASE/rover/commit/451f7eb4ab031fe82dc0d7b19e2a7377122a2e5f))
* roverd log lines ([97adc00](https://github.com/VU-ASE/rover/commit/97adc0044f551ce2aa9ea1e7db9beeae24e43ab7))
* roverd not propagating stream outputs to transceiver if no consumer ([0818f38](https://github.com/VU-ASE/rover/commit/0818f38c0b4353bb5078a3c8b904be8a9e5f3c38))
* roverd shows all exit codes now ([2de0226](https://github.com/VU-ASE/rover/commit/2de022691173795eac6a63eec9787d554f466958))
* roverd shutsdown pipeline on termination ([83ea1dc](https://github.com/VU-ASE/rover/commit/83ea1dce347970d9a265472b62a18031177b46eb))
* roverd unclosed delimiter ([17d8e95](https://github.com/VU-ASE/rover/commit/17d8e95a9b322a205ffe46c5dd367459f59c9aef))
* run open-api generator as devuser ([e8e84c8](https://github.com/VU-ASE/rover/commit/e8e84c85578d578299798f469782f04e9acff0f0))
* service logs cut-off point makes them unreadable ([412a8ee](https://github.com/VU-ASE/rover/commit/412a8ee5eb7520daee8fbef481b0271eadab9470))
* service yamls can now specify a bash command for run ([6022928](https://github.com/VU-ASE/rover/commit/60229285d294b95698faef9958704e3d8c81ff5e))
* show correct fix for updating roverd on mismatch ([a49144f](https://github.com/VU-ASE/rover/commit/a49144feec173a050d84a6bfee482bcb7fe7afc3))
* show correct panels for selected services ([5ba569e](https://github.com/VU-ASE/rover/commit/5ba569e06f205686fd5baa2dd7ff0ae765d06bb7))
* show debug toggle correctly ([3ea63f8](https://github.com/VU-ASE/rover/commit/3ea63f8a79ee4fe7ca0836407fb96d6f501fd94d))
* show logs of last run when selected ([0cca2da](https://github.com/VU-ASE/rover/commit/0cca2dab6e762f6fb8916ed3cce3175f17262ec0))
* show pipeline edges accurately ([ff53e34](https://github.com/VU-ASE/rover/commit/ff53e3424e6c8644bd943c3d181c080cf69f6a74))
* show service crash exit code on transition from started to stopped ([e43cb26](https://github.com/VU-ASE/rover/commit/e43cb260ef1df98869502afda7731fe188257beb))
* sort services alphabetically. Always show vu-ase first ([fb6a58b](https://github.com/VU-ASE/rover/commit/fb6a58b9419af46bb1b06792e548bdc9431fb6da))
* specifying value types in service.yaml is required ([e8757a3](https://github.com/VU-ASE/rover/commit/e8757a3edeb67edcccb0cd5d6909d15879d6f886))
* testing ci again ([573bb0f](https://github.com/VU-ASE/rover/commit/573bb0f5aafb6ca7b18208a4c71b1a14724de484))
* transceiver output address should use `*` instead of `localhost` ([19264da](https://github.com/VU-ASE/rover/commit/19264dac0d2c18e4ff8485889e4f9dd0dd0fddab))
* update checking for default pipeline ([3f67220](https://github.com/VU-ASE/rover/commit/3f672204d7faaeae6ec4a3d2ec141ab52af4dc55))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))
* use correct version of action/artifact ([300beda](https://github.com/VU-ASE/rover/commit/300beda2031241877d4101279e5f9ec5cf8dcd07))
* use docker buildx for builds ([aa4893b](https://github.com/VU-ASE/rover/commit/aa4893b9e97492b82198164711f7cca2e534f094))
* use dynamic env variables for configuration ([25d5c88](https://github.com/VU-ASE/rover/commit/25d5c88b217c9040aea12686a4b08b05443a8e03))
* use lowercase for GHCR tag ([e62a583](https://github.com/VU-ASE/rover/commit/e62a58302e17c7ca9739349b60fc654cce7f7da3))
* use mdns for tests ([bb59b70](https://github.com/VU-ASE/rover/commit/bb59b70cbbbc7e8a31fdf4e93a85f63346c22e5f))
* using mv instead of os.Rename ([696ecc4](https://github.com/VU-ASE/rover/commit/696ecc440505fb39c3ec5cf814b55786b9fb5bf5))
* version mismatch checking ([4ba63b9](https://github.com/VU-ASE/rover/commit/4ba63b901f88055e68db2a7234ca17659588798c))
* version update script now accepts "v" prefix ([362c24c](https://github.com/VU-ASE/rover/commit/362c24c0b3264ce7d3aef7fdfabbfb6bea72d45a))
* version utility and correct comparison ([9251512](https://github.com/VU-ASE/rover/commit/92515124a450c2edd2ad8aed10a20b7b42ba4b9c))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))
* release 0.0.1 ([abd806e](https://github.com/VU-ASE/rover/commit/abd806e59353ee5ea9da82a80612ff9ece653285))
* release 0.0.5 ([aa2ed11](https://github.com/VU-ASE/rover/commit/aa2ed11e18453f8a6edc3f82cd503e38288e7060))
* release 1.0.0 ([b1c1497](https://github.com/VU-ASE/rover/commit/b1c14978496488dac92350a6c75cb27156cd017a))
* release 1.0.0 ([55d09fd](https://github.com/VU-ASE/rover/commit/55d09fd29d0bc18c528ce63af45a2c10c3428f8d))
* release 1.0.1 ([ea40d3b](https://github.com/VU-ASE/rover/commit/ea40d3bcfe9ac8d3147e4b2873d05a1b305d348f))

## [1.0.0](https://github.com/VU-ASE/rover/compare/v1.0.0...v1.0.0) (2025-04-01)


### ⚠ BREAKING CHANGES

* release rover v1.0.0

### Features

* add alias "rover" to roverctl install script ([1ec8389](https://github.com/VU-ASE/rover/commit/1ec83893a0a7458aef2c8b41791c548ffda961df))
* add incompatibility warning to all endpoints that need it ([126849c](https://github.com/VU-ASE/rover/commit/126849cd89e743d089f2765f7261df24c30fd0eb))
* add installation location to service info view ([1a0ca7c](https://github.com/VU-ASE/rover/commit/1a0ca7c58c0a9b58ef3a0073babc92470ce99fe1))
* add passthrough back ([e07dc70](https://github.com/VU-ASE/rover/commit/e07dc7073c6b63dd9ed1c9894a11df2127e7d23c))
* add shortcuts item in navbar ([5aea4d3](https://github.com/VU-ASE/rover/commit/5aea4d3c7ec4458cd4aa6feb79e62398d5c1d860))
* added /emergency endpoint ([55df2dc](https://github.com/VU-ASE/rover/commit/55df2dcf7d51e36a6bec730cd36b2324de9cd45b))
* added Dockerfile for roverctl-web ([ec9733b](https://github.com/VU-ASE/rover/commit/ec9733b2f507d03d0ec405e6bcc2d82519406e6b))
* added manage script ([e1b1429](https://github.com/VU-ASE/rover/commit/e1b142989b1b0c0dfcaa55f0a496da56031288e4))
* added passthrough and debugging functionality ([90975ef](https://github.com/VU-ASE/rover/commit/90975ef34e1acf2ffc1f9ec740e18b600df7d561))
* api adaptation ([d36262f](https://github.com/VU-ASE/rover/commit/d36262f0e0d336e08cc1a3a6e2484ad4e52c2431))
* api clean up and addition of alias error (BBC) ([1d685a6](https://github.com/VU-ASE/rover/commit/1d685a6578ddc0b9c3b48fa9ffeb0abecfefad22))
* app config and overview page ([1c76d8c](https://github.com/VU-ASE/rover/commit/1c76d8cd42abe86a77e5133806ac97921a4da5d5))
* author command ([6ead41c](https://github.com/VU-ASE/rover/commit/6ead41cdbfb58ab0ce484cb627ae01bae0223da1))
* build command ([ea610eb](https://github.com/VU-ASE/rover/commit/ea610eb0171be57fbfd41ca2991bc6a1d00dedea))
* calibration command for actuator service ([a549e94](https://github.com/VU-ASE/rover/commit/a549e9499f9b629e89b40ea6db1fc4e1e6f025c4))
* CLI approach to roverctl ([bb78304](https://github.com/VU-ASE/rover/commit/bb78304cf2c6fcda4b2ce95bba51fcd183ee6ce3))
* dangers page ([a296ad3](https://github.com/VU-ASE/rover/commit/a296ad3ccb2d682d7a881c057d5a60aec36a9888))
* default ASE pipeline updates and installs ([0abcdb0](https://github.com/VU-ASE/rover/commit/0abcdb0021d0c2c4158a52e479684615864518af))
* dialog and upgrad/downgrade roverctl on version mismatch ([5d25e52](https://github.com/VU-ASE/rover/commit/5d25e52df3a75dfad549aca7e929f39d0755aa6f))
* emergency command ([2b883cd](https://github.com/VU-ASE/rover/commit/2b883cd7ded78acf211950161ec854884722badd))
* enable and install debug logic from roverctl-web ([a760e30](https://github.com/VU-ASE/rover/commit/a760e30d6991e0322399aeb4d65beca78a06dcc6))
* error types for the configuration endpoint ([016a94b](https://github.com/VU-ASE/rover/commit/016a94bd55a8643d300bf94d1c88c54bdb60512b))
* github actions workflows to build and push docker images ([5c4b035](https://github.com/VU-ASE/rover/commit/5c4b035329c9998331228113913a28fcc72b7552))
* group test scripts ([d073bfd](https://github.com/VU-ASE/rover/commit/d073bfdecdd198ae74f7b697144ed7dee806826c))
* helpful toasts and warning messages when pipeline starts/stops ([27074f6](https://github.com/VU-ASE/rover/commit/27074f67ca459497ca9105b6ec2c419f70a2d49b))
* implemented in place update API ([44a8c73](https://github.com/VU-ASE/rover/commit/44a8c73ffd2ad93d22056335e78b3a8ff4a5da39))
* improved error printing with multiline support ([8c64c2c](https://github.com/VU-ASE/rover/commit/8c64c2ce5443774ab1571658d5f57e932cb426b1))
* improved error reporting on upload failure ([b657f17](https://github.com/VU-ASE/rover/commit/b657f17d60bdaf464a2e998ddbb1541fead96141))
* info command ([478704e](https://github.com/VU-ASE/rover/commit/478704ec6fa6d5998f16da28bf9116df0c3dac45))
* info target ([0bc9cb0](https://github.com/VU-ASE/rover/commit/0bc9cb0ecfd0d6ffe2aac31c12717839d7ede73f))
* install command ([955fea9](https://github.com/VU-ASE/rover/commit/955fea970534b51d0e8e31abef62600f776475c1))
* install script that allows for installing a specific version ([f5bfcaf](https://github.com/VU-ASE/rover/commit/f5bfcaf2bada29e9e8732b61d79c93a53c3331d9))
* install script with symlink instead of alias ([22ccd67](https://github.com/VU-ASE/rover/commit/22ccd6786c112cbde2979064db8ae6ef89fd5c46))
* keyboard shortcuts for emergency, stopping pipelin and debug pause ([ad50f10](https://github.com/VU-ASE/rover/commit/ad50f1014705f1631c973928ff06922c74733514))
* link source code and docs for official services ([3ac9131](https://github.com/VU-ASE/rover/commit/3ac9131e4fce750e5475ec18167cf9feb83109c1))
* logs endpoint ([f1b3077](https://github.com/VU-ASE/rover/commit/f1b3077fb46bcc396dc9a7c1d32c96b8a66537d7))
* manage page (wip) and status overlay ([e9132ed](https://github.com/VU-ASE/rover/commit/e9132ed41535b49fbc10353c023d6935edf82640))
* merge passthrough into roverctl, debugging now works ([8f6c88f](https://github.com/VU-ASE/rover/commit/8f6c88f43323856971c7f6f129c47a0d23ba3a1e))
* pin version dependency between roverctl-roverd and warn if broken ([c8e98dc](https://github.com/VU-ASE/rover/commit/c8e98dcc4b89d079625197e91685f5b085be4bfd))
* pipeline enable/disable endpoints ([148fb7e](https://github.com/VU-ASE/rover/commit/148fb7eb16da42ae3af8ce568a767419260d9704))
* pipeline endpoint ([d0beca5](https://github.com/VU-ASE/rover/commit/d0beca50b2d01d0cedaada7833edd6a8bcba1053))
* release rover v1.0.0 ([292c229](https://github.com/VU-ASE/rover/commit/292c22939aeb4b8592c2dbe8ac1566c350d83b1a))
* reset calibration values on trimming ([ec8eb15](https://github.com/VU-ASE/rover/commit/ec8eb15b7efc09a2b245a1424f21fb4bb51beb9d))
* roverctl service init endpoint ([b8a9027](https://github.com/VU-ASE/rover/commit/b8a90274bc33f681fc5eaf97cb8adf04b828122f))
* roverctl-web startup through docker ([e3ca9a5](https://github.com/VU-ASE/rover/commit/e3ca9a54be65233a5f7a8f8fa8b915c9d042b724))
* service delete endpoint ([185362f](https://github.com/VU-ASE/rover/commit/185362f0f91d3af1cac2b0abe58e4741260740f9))
* service information ([c0119a5](https://github.com/VU-ASE/rover/commit/c0119a5c4715cb39d86fdecdf856ce3d48127595))
* service installation ([246a452](https://github.com/VU-ASE/rover/commit/246a452e0ae9b4d8489b1f852d1a425c7137e895))
* service selection toggles and ordening ([6ad02f4](https://github.com/VU-ASE/rover/commit/6ad02f49222b35caec422dc1e2c486276a50aa65))
* service_as feature ([782ab78](https://github.com/VU-ASE/rover/commit/782ab7874c1eb6eec0e58ac73283f2f9ab940a5b))
* services endpoint ([a5e11a4](https://github.com/VU-ASE/rover/commit/a5e11a46ca7d8c5a20345a356f32943e86ba50f0))
* shorthand aliasses for most commands ([3c270ce](https://github.com/VU-ASE/rover/commit/3c270ce7e0fbff41f1273a2cedd927f88d24fb99))
* show transceiver in roverctl-web ([be2cb92](https://github.com/VU-ASE/rover/commit/be2cb9274d7bed0c2d4e890b33018c8f82fbd652))
* show which service crashed the pipeline ([8591295](https://github.com/VU-ASE/rover/commit/85912954080cdb6654596702b195fb1c2b064ded))
* shutdown command ([4a124b7](https://github.com/VU-ASE/rover/commit/4a124b7ddec97a31e733b064745a4374e4cd71fc))
* start roverctl-web container ([9329015](https://github.com/VU-ASE/rover/commit/9329015b9cb86a8db86bc756052327fdf170785e))
* starting, stopping and modifying pipelines ([e54a176](https://github.com/VU-ASE/rover/commit/e54a1762fdc686ec0236874c1169ea9c20014bff))
* unified toasts ([ded557b](https://github.com/VU-ASE/rover/commit/ded557be905b07b688a2fe9f6098a028543ec8bb))
* update endpoints ([e1422b4](https://github.com/VU-ASE/rover/commit/e1422b40ed2e25e98eb094bc088154d896a28caf))
* update target ([0f3cc2a](https://github.com/VU-ASE/rover/commit/0f3cc2a340b5a78a3c04fc363941608fa9622613))
* upgrade roverctl to go 1.22.12 ([6a19635](https://github.com/VU-ASE/rover/commit/6a1963502affb37a09328151d1c55622ffb14d0f))
* use mDNS to resolve hostnames, allow proxy IP override ([2cb436e](https://github.com/VU-ASE/rover/commit/2cb436e2038e2ef1e5d722913f0d794ce9d60bfe))
* warn users if roverctl mismatches roverd and install appr roverctl ([868e3c1](https://github.com/VU-ASE/rover/commit/868e3c1db16c7dc0955121ae022ded73ff152f0a))


### Bug Fixes

* added service-template for cpp ([cf5fd0c](https://github.com/VU-ASE/rover/commit/cf5fd0c04b646dfef25138bbbb54220c03a2722e))
* added smoother steering to testing scripts ([681d881](https://github.com/VU-ASE/rover/commit/681d881c2147039cd8d2954098833fd0110bc080))
* auto-update roverd in test ([9b53863](https://github.com/VU-ASE/rover/commit/9b53863f55a2baf712b25f46099b1d2bbd7b7ff3))
* broken links ([7a4d2aa](https://github.com/VU-ASE/rover/commit/7a4d2aa8ee9bc87d5f06ea3c1f55a9f324cc57a7))
* build to ./bin ([dba0f2b](https://github.com/VU-ASE/rover/commit/dba0f2b1cfd4deaa257beb7c0bbe241b02879f5d))
* ci testing ([269b7a3](https://github.com/VU-ASE/rover/commit/269b7a3b9ece651b79b22145db9eccb1cf018fa9))
* consistent style for services ([a9f5fbe](https://github.com/VU-ASE/rover/commit/a9f5fbef311b50ce30ca3891a9c5591ec3c12d91))
* correct link if upload fails ([d229edb](https://github.com/VU-ASE/rover/commit/d229edbeca49c5684dd881e97dc9d02bf9cbb9fc))
* correct timing of service error messages to actual moment of crash ([68f158f](https://github.com/VU-ASE/rover/commit/68f158f9404125399fc18499d855f04f6ab054f0))
* create openapi directory if it does not exist ([1f881b8](https://github.com/VU-ASE/rover/commit/1f881b8846f38ceea597ab31be031ae7c60ecf96))
* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* disable transceiver service when debug mode is not enabled ([46f41e5](https://github.com/VU-ASE/rover/commit/46f41e5f75025b3dd3515852eb42f69f88afccea))
* display warning if author is not set in info command ([b2c5eaa](https://github.com/VU-ASE/rover/commit/b2c5eaa5e866c87066cdd2eddc1aecd6f680d493))
* do not build roverd on self-hosted ([af62651](https://github.com/VU-ASE/rover/commit/af626516335e3f1c2a5682aaa8fb54b5f72deb5c))
* do not error on existing connection being reused ([6e357bf](https://github.com/VU-ASE/rover/commit/6e357bfd6dc7c008d09f0717992a3619f7f2240c))
* do not publish openapi files ([f7f7ff0](https://github.com/VU-ASE/rover/commit/f7f7ff039253c5ebfc893fa0d53b5daad304968d))
* do not reset calibration value for normal pipeline testing ([be21d46](https://github.com/VU-ASE/rover/commit/be21d462d992f50a90e93b473417e6a0f026b3f3))
* do not show "shortcuts" tab as active on debugpage ([822b084](https://github.com/VU-ASE/rover/commit/822b084f1a15a00c34a12a162d8c618e61bc5c65))
* do NOT update roverd in the test script ([b5f1cb0](https://github.com/VU-ASE/rover/commit/b5f1cb0ea34f527117f6aee6a8d071002786c52d))
* docs and install script for roverctl ([a069205](https://github.com/VU-ASE/rover/commit/a069205b726187a5ad7ad99c8f00a75f7569ce9f))
* emergency stop ([2b7bdf8](https://github.com/VU-ASE/rover/commit/2b7bdf85012d4b5e736a5bf3543a7f5856bc9855))
* error on roverctl install.sh if release does not exist or no binary ([1880064](https://github.com/VU-ASE/rover/commit/1880064e54514a502ae07bd21f690b8d953a7eea))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* extract artifact downloads first ([37762b7](https://github.com/VU-ASE/rover/commit/37762b77e75becc3a8fb960e83ddd72f47305729))
* extract from correct location ([d4106ba](https://github.com/VU-ASE/rover/commit/d4106ba17d7e8ec4d7fea3f3827aded9d4a25d0a))
* gitignore was missing readme ([44c64ff](https://github.com/VU-ASE/rover/commit/44c64ffed3c5d317704599bfb95af5ace86af280))
* hard-coded version ([058573d](https://github.com/VU-ASE/rover/commit/058573d9cf67a814ef554712b4ad610eaabcb0b4))
* hide updater screen when all services are installed ([f4dce16](https://github.com/VU-ASE/rover/commit/f4dce16e2ae91666336fa361d6c2d15db813c380))
* image push only succeeds on --verbose flag, port checking ([b0d5998](https://github.com/VU-ASE/rover/commit/b0d599895c6cc1f3365eca1aee7ceadbadb6e13e))
* improved upload description and example usage ([248002c](https://github.com/VU-ASE/rover/commit/248002c3d1c96494a91577621a270ea0a3a68cea))
* incorrect number of arguments for service init script ([49981a0](https://github.com/VU-ASE/rover/commit/49981a094084063c1f584caf0f3df2a25e13aeca))
* infinite redirect ([0f63b01](https://github.com/VU-ASE/rover/commit/0f63b019402c498c496f4cec1bbf9f483de3065a))
* load list of services from roverd in debug window ([a8099bb](https://github.com/VU-ASE/rover/commit/a8099bb929543b3dfde14b06bc61115c880f3ceb))
* make roverctl compatible with new API spec ([6eb675c](https://github.com/VU-ASE/rover/commit/6eb675c87be7b49d31ca3405906767df359405f3))
* merge service name and service "as" name in tuning overview ([3675c0e](https://github.com/VU-ASE/rover/commit/3675c0e5159f0e503c8ff65a596c5f757263a9c5))
* more helpful messages to enable debug mode ([106e3ea](https://github.com/VU-ASE/rover/commit/106e3ea929bc7943948e77f07074171e4a31ad55))
* more lenient timeouts for roverd downloads ([a89b302](https://github.com/VU-ASE/rover/commit/a89b302fd97b48ce176d1f27becee737a67ac803))
* move github.com/VU-ASE/roverctl to github.com/VU-ASE/rover/roverctl ([a35e0dc](https://github.com/VU-ASE/rover/commit/a35e0dc12895750545256432284054e560e9eb65))
* navbar type error ([d66c700](https://github.com/VU-ASE/rover/commit/d66c70009adc0a222eba5967af53d168de2f44af))
* only show services when the pipeline is started (in debug mode) ([e22facf](https://github.com/VU-ASE/rover/commit/e22facf86deb8c7b305238dcc704b6aab7fc741b))
* open-api-gen debugging ([013bce9](https://github.com/VU-ASE/rover/commit/013bce9f4a0838ccf91ddb53e429d27cb7132178))
* permission issues for building ([8f82d02](https://github.com/VU-ASE/rover/commit/8f82d02b4b4bcd6a5930c5faa8fd504042a69f0b))
* permissions on github actions ([0b7225a](https://github.com/VU-ASE/rover/commit/0b7225ae596dd3d1e1a29a46ad26b2f6e256d9dd))
* pipeline not loading after viewing debug screen ([67787c4](https://github.com/VU-ASE/rover/commit/67787c480eb8d8c92381433008e21b2297b5b6eb))
* point install scripts to rover repository (was roverctl) ([b139324](https://github.com/VU-ASE/rover/commit/b139324d303d27c88f2f8d91d0aaf1b78a5f714a))
* poll pipeline status ([a02bbf6](https://github.com/VU-ASE/rover/commit/a02bbf698dbd8888b4f8efcf82dfb2fdb7eca9a8))
* proxy server deadlock on repeated transceiver crashes ([577a1a2](https://github.com/VU-ASE/rover/commit/577a1a29775debc7e75526abbc9f799e9b3db0ec))
* public link for schematic in docs ([b26decf](https://github.com/VU-ASE/rover/commit/b26decf57eb08b15e0fd4c2c04e81d1f9b3a1fdd))
* release assets on all success ([4448019](https://github.com/VU-ASE/rover/commit/44480199fcac3cc9c81e317ef7b34c270641ec6a))
* remove unused github action jobs ([04ef005](https://github.com/VU-ASE/rover/commit/04ef0051fa635a52e6fbc521c6ca915b58c1423e))
* remove unused view ([31f91cd](https://github.com/VU-ASE/rover/commit/31f91cd1bd482a2a1bb05af107ebb19b153d1de5))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* rename passthrough to appropriate folder ([5c1d41e](https://github.com/VU-ASE/rover/commit/5c1d41e8f3a1bb3be425841ede4c9a80fe2c2861))
* retain debug toggle information and don't show unusable tuning ([9d78947](https://github.com/VU-ASE/rover/commit/9d789474196bc613cea17b1da0a775e8669e9d12))
* returning service_as optimistically for all endpoints ([a019fe1](https://github.com/VU-ASE/rover/commit/a019fe1bee5a0a5e591911b389ee599b647e0440))
* roverctl --force warning message ([5384c70](https://github.com/VU-ASE/rover/commit/5384c704cac253524bc7d433d4b260602e1c3632))
* roverctl install.sh, verify downloaded binary is valid ([381a5ed](https://github.com/VU-ASE/rover/commit/381a5ed2e809710a882a65e4affd3fb658ecb2e2))
* roverctl release 1.0.0 ([ec304d4](https://github.com/VU-ASE/rover/commit/ec304d46cf0205bed12551c316e67ad69d277c0b))
* roverctl replace FQN types ([f6a9c44](https://github.com/VU-ASE/rover/commit/f6a9c44dee576b4ba63f3e44ca5a07704ead245b))
* roverctl-roverd update API endpoint not working ([9fd776f](https://github.com/VU-ASE/rover/commit/9fd776f4616d2c66e59a66cb0c2b554620a72522))
* roverctl, build openapi files in Dockerfile ([af01eb6](https://github.com/VU-ASE/rover/commit/af01eb620b565518860cb84d5665595788b0de71))
* roverd early check for root ([451f7eb](https://github.com/VU-ASE/rover/commit/451f7eb4ab031fe82dc0d7b19e2a7377122a2e5f))
* roverd log lines ([97adc00](https://github.com/VU-ASE/rover/commit/97adc0044f551ce2aa9ea1e7db9beeae24e43ab7))
* roverd not propagating stream outputs to transceiver if no consumer ([0818f38](https://github.com/VU-ASE/rover/commit/0818f38c0b4353bb5078a3c8b904be8a9e5f3c38))
* roverd shows all exit codes now ([2de0226](https://github.com/VU-ASE/rover/commit/2de022691173795eac6a63eec9787d554f466958))
* roverd shutsdown pipeline on termination ([83ea1dc](https://github.com/VU-ASE/rover/commit/83ea1dce347970d9a265472b62a18031177b46eb))
* roverd unclosed delimiter ([17d8e95](https://github.com/VU-ASE/rover/commit/17d8e95a9b322a205ffe46c5dd367459f59c9aef))
* run open-api generator as devuser ([e8e84c8](https://github.com/VU-ASE/rover/commit/e8e84c85578d578299798f469782f04e9acff0f0))
* service logs cut-off point makes them unreadable ([412a8ee](https://github.com/VU-ASE/rover/commit/412a8ee5eb7520daee8fbef481b0271eadab9470))
* service yamls can now specify a bash command for run ([6022928](https://github.com/VU-ASE/rover/commit/60229285d294b95698faef9958704e3d8c81ff5e))
* show correct fix for updating roverd on mismatch ([a49144f](https://github.com/VU-ASE/rover/commit/a49144feec173a050d84a6bfee482bcb7fe7afc3))
* show correct panels for selected services ([5ba569e](https://github.com/VU-ASE/rover/commit/5ba569e06f205686fd5baa2dd7ff0ae765d06bb7))
* show debug toggle correctly ([3ea63f8](https://github.com/VU-ASE/rover/commit/3ea63f8a79ee4fe7ca0836407fb96d6f501fd94d))
* show logs of last run when selected ([0cca2da](https://github.com/VU-ASE/rover/commit/0cca2dab6e762f6fb8916ed3cce3175f17262ec0))
* show pipeline edges accurately ([ff53e34](https://github.com/VU-ASE/rover/commit/ff53e3424e6c8644bd943c3d181c080cf69f6a74))
* show service crash exit code on transition from started to stopped ([e43cb26](https://github.com/VU-ASE/rover/commit/e43cb260ef1df98869502afda7731fe188257beb))
* sort services alphabetically. Always show vu-ase first ([fb6a58b](https://github.com/VU-ASE/rover/commit/fb6a58b9419af46bb1b06792e548bdc9431fb6da))
* specifying value types in service.yaml is required ([e8757a3](https://github.com/VU-ASE/rover/commit/e8757a3edeb67edcccb0cd5d6909d15879d6f886))
* testing ci again ([573bb0f](https://github.com/VU-ASE/rover/commit/573bb0f5aafb6ca7b18208a4c71b1a14724de484))
* transceiver output address should use `*` instead of `localhost` ([19264da](https://github.com/VU-ASE/rover/commit/19264dac0d2c18e4ff8485889e4f9dd0dd0fddab))
* update checking for default pipeline ([3f67220](https://github.com/VU-ASE/rover/commit/3f672204d7faaeae6ec4a3d2ec141ab52af4dc55))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))
* use correct version of action/artifact ([300beda](https://github.com/VU-ASE/rover/commit/300beda2031241877d4101279e5f9ec5cf8dcd07))
* use docker buildx for builds ([aa4893b](https://github.com/VU-ASE/rover/commit/aa4893b9e97492b82198164711f7cca2e534f094))
* use dynamic env variables for configuration ([25d5c88](https://github.com/VU-ASE/rover/commit/25d5c88b217c9040aea12686a4b08b05443a8e03))
* use lowercase for GHCR tag ([e62a583](https://github.com/VU-ASE/rover/commit/e62a58302e17c7ca9739349b60fc654cce7f7da3))
* use mdns for tests ([bb59b70](https://github.com/VU-ASE/rover/commit/bb59b70cbbbc7e8a31fdf4e93a85f63346c22e5f))
* using mv instead of os.Rename ([696ecc4](https://github.com/VU-ASE/rover/commit/696ecc440505fb39c3ec5cf814b55786b9fb5bf5))
* version mismatch checking ([4ba63b9](https://github.com/VU-ASE/rover/commit/4ba63b901f88055e68db2a7234ca17659588798c))
* version update script now accepts "v" prefix ([362c24c](https://github.com/VU-ASE/rover/commit/362c24c0b3264ce7d3aef7fdfabbfb6bea72d45a))
* version utility and correct comparison ([9251512](https://github.com/VU-ASE/rover/commit/92515124a450c2edd2ad8aed10a20b7b42ba4b9c))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))
* release 0.0.1 ([abd806e](https://github.com/VU-ASE/rover/commit/abd806e59353ee5ea9da82a80612ff9ece653285))
* release 0.0.5 ([aa2ed11](https://github.com/VU-ASE/rover/commit/aa2ed11e18453f8a6edc3f82cd503e38288e7060))
* release 1.0.0 ([b1c1497](https://github.com/VU-ASE/rover/commit/b1c14978496488dac92350a6c75cb27156cd017a))
* release 1.0.0 ([55d09fd](https://github.com/VU-ASE/rover/commit/55d09fd29d0bc18c528ce63af45a2c10c3428f8d))

## [1.0.0](https://github.com/VU-ASE/rover/compare/v1.0.0...v1.0.0) (2025-04-01)


### ⚠ BREAKING CHANGES

* release rover v1.0.0

### Features

* add alias "rover" to roverctl install script ([1ec8389](https://github.com/VU-ASE/rover/commit/1ec83893a0a7458aef2c8b41791c548ffda961df))
* add incompatibility warning to all endpoints that need it ([126849c](https://github.com/VU-ASE/rover/commit/126849cd89e743d089f2765f7261df24c30fd0eb))
* add installation location to service info view ([1a0ca7c](https://github.com/VU-ASE/rover/commit/1a0ca7c58c0a9b58ef3a0073babc92470ce99fe1))
* add passthrough back ([e07dc70](https://github.com/VU-ASE/rover/commit/e07dc7073c6b63dd9ed1c9894a11df2127e7d23c))
* add shortcuts item in navbar ([5aea4d3](https://github.com/VU-ASE/rover/commit/5aea4d3c7ec4458cd4aa6feb79e62398d5c1d860))
* added /emergency endpoint ([55df2dc](https://github.com/VU-ASE/rover/commit/55df2dcf7d51e36a6bec730cd36b2324de9cd45b))
* added Dockerfile for roverctl-web ([ec9733b](https://github.com/VU-ASE/rover/commit/ec9733b2f507d03d0ec405e6bcc2d82519406e6b))
* added manage script ([e1b1429](https://github.com/VU-ASE/rover/commit/e1b142989b1b0c0dfcaa55f0a496da56031288e4))
* added passthrough and debugging functionality ([90975ef](https://github.com/VU-ASE/rover/commit/90975ef34e1acf2ffc1f9ec740e18b600df7d561))
* api adaptation ([d36262f](https://github.com/VU-ASE/rover/commit/d36262f0e0d336e08cc1a3a6e2484ad4e52c2431))
* api clean up and addition of alias error (BBC) ([1d685a6](https://github.com/VU-ASE/rover/commit/1d685a6578ddc0b9c3b48fa9ffeb0abecfefad22))
* app config and overview page ([1c76d8c](https://github.com/VU-ASE/rover/commit/1c76d8cd42abe86a77e5133806ac97921a4da5d5))
* author command ([6ead41c](https://github.com/VU-ASE/rover/commit/6ead41cdbfb58ab0ce484cb627ae01bae0223da1))
* build command ([ea610eb](https://github.com/VU-ASE/rover/commit/ea610eb0171be57fbfd41ca2991bc6a1d00dedea))
* calibration command for actuator service ([a549e94](https://github.com/VU-ASE/rover/commit/a549e9499f9b629e89b40ea6db1fc4e1e6f025c4))
* CLI approach to roverctl ([bb78304](https://github.com/VU-ASE/rover/commit/bb78304cf2c6fcda4b2ce95bba51fcd183ee6ce3))
* dangers page ([a296ad3](https://github.com/VU-ASE/rover/commit/a296ad3ccb2d682d7a881c057d5a60aec36a9888))
* default ASE pipeline updates and installs ([0abcdb0](https://github.com/VU-ASE/rover/commit/0abcdb0021d0c2c4158a52e479684615864518af))
* dialog and upgrad/downgrade roverctl on version mismatch ([5d25e52](https://github.com/VU-ASE/rover/commit/5d25e52df3a75dfad549aca7e929f39d0755aa6f))
* emergency command ([2b883cd](https://github.com/VU-ASE/rover/commit/2b883cd7ded78acf211950161ec854884722badd))
* enable and install debug logic from roverctl-web ([a760e30](https://github.com/VU-ASE/rover/commit/a760e30d6991e0322399aeb4d65beca78a06dcc6))
* error types for the configuration endpoint ([016a94b](https://github.com/VU-ASE/rover/commit/016a94bd55a8643d300bf94d1c88c54bdb60512b))
* github actions workflows to build and push docker images ([5c4b035](https://github.com/VU-ASE/rover/commit/5c4b035329c9998331228113913a28fcc72b7552))
* group test scripts ([d073bfd](https://github.com/VU-ASE/rover/commit/d073bfdecdd198ae74f7b697144ed7dee806826c))
* helpful toasts and warning messages when pipeline starts/stops ([27074f6](https://github.com/VU-ASE/rover/commit/27074f67ca459497ca9105b6ec2c419f70a2d49b))
* implemented in place update API ([44a8c73](https://github.com/VU-ASE/rover/commit/44a8c73ffd2ad93d22056335e78b3a8ff4a5da39))
* improved error printing with multiline support ([8c64c2c](https://github.com/VU-ASE/rover/commit/8c64c2ce5443774ab1571658d5f57e932cb426b1))
* improved error reporting on upload failure ([b657f17](https://github.com/VU-ASE/rover/commit/b657f17d60bdaf464a2e998ddbb1541fead96141))
* info command ([478704e](https://github.com/VU-ASE/rover/commit/478704ec6fa6d5998f16da28bf9116df0c3dac45))
* info target ([0bc9cb0](https://github.com/VU-ASE/rover/commit/0bc9cb0ecfd0d6ffe2aac31c12717839d7ede73f))
* install command ([955fea9](https://github.com/VU-ASE/rover/commit/955fea970534b51d0e8e31abef62600f776475c1))
* install script that allows for installing a specific version ([f5bfcaf](https://github.com/VU-ASE/rover/commit/f5bfcaf2bada29e9e8732b61d79c93a53c3331d9))
* install script with symlink instead of alias ([22ccd67](https://github.com/VU-ASE/rover/commit/22ccd6786c112cbde2979064db8ae6ef89fd5c46))
* keyboard shortcuts for emergency, stopping pipelin and debug pause ([ad50f10](https://github.com/VU-ASE/rover/commit/ad50f1014705f1631c973928ff06922c74733514))
* link source code and docs for official services ([3ac9131](https://github.com/VU-ASE/rover/commit/3ac9131e4fce750e5475ec18167cf9feb83109c1))
* logs endpoint ([f1b3077](https://github.com/VU-ASE/rover/commit/f1b3077fb46bcc396dc9a7c1d32c96b8a66537d7))
* manage page (wip) and status overlay ([e9132ed](https://github.com/VU-ASE/rover/commit/e9132ed41535b49fbc10353c023d6935edf82640))
* merge passthrough into roverctl, debugging now works ([8f6c88f](https://github.com/VU-ASE/rover/commit/8f6c88f43323856971c7f6f129c47a0d23ba3a1e))
* pin version dependency between roverctl-roverd and warn if broken ([c8e98dc](https://github.com/VU-ASE/rover/commit/c8e98dcc4b89d079625197e91685f5b085be4bfd))
* pipeline enable/disable endpoints ([148fb7e](https://github.com/VU-ASE/rover/commit/148fb7eb16da42ae3af8ce568a767419260d9704))
* pipeline endpoint ([d0beca5](https://github.com/VU-ASE/rover/commit/d0beca50b2d01d0cedaada7833edd6a8bcba1053))
* release rover v1.0.0 ([292c229](https://github.com/VU-ASE/rover/commit/292c22939aeb4b8592c2dbe8ac1566c350d83b1a))
* reset calibration values on trimming ([ec8eb15](https://github.com/VU-ASE/rover/commit/ec8eb15b7efc09a2b245a1424f21fb4bb51beb9d))
* roverctl service init endpoint ([b8a9027](https://github.com/VU-ASE/rover/commit/b8a90274bc33f681fc5eaf97cb8adf04b828122f))
* roverctl-web startup through docker ([e3ca9a5](https://github.com/VU-ASE/rover/commit/e3ca9a54be65233a5f7a8f8fa8b915c9d042b724))
* service delete endpoint ([185362f](https://github.com/VU-ASE/rover/commit/185362f0f91d3af1cac2b0abe58e4741260740f9))
* service information ([c0119a5](https://github.com/VU-ASE/rover/commit/c0119a5c4715cb39d86fdecdf856ce3d48127595))
* service installation ([246a452](https://github.com/VU-ASE/rover/commit/246a452e0ae9b4d8489b1f852d1a425c7137e895))
* service selection toggles and ordening ([6ad02f4](https://github.com/VU-ASE/rover/commit/6ad02f49222b35caec422dc1e2c486276a50aa65))
* service_as feature ([782ab78](https://github.com/VU-ASE/rover/commit/782ab7874c1eb6eec0e58ac73283f2f9ab940a5b))
* services endpoint ([a5e11a4](https://github.com/VU-ASE/rover/commit/a5e11a46ca7d8c5a20345a356f32943e86ba50f0))
* shorthand aliasses for most commands ([3c270ce](https://github.com/VU-ASE/rover/commit/3c270ce7e0fbff41f1273a2cedd927f88d24fb99))
* show transceiver in roverctl-web ([be2cb92](https://github.com/VU-ASE/rover/commit/be2cb9274d7bed0c2d4e890b33018c8f82fbd652))
* show which service crashed the pipeline ([8591295](https://github.com/VU-ASE/rover/commit/85912954080cdb6654596702b195fb1c2b064ded))
* shutdown command ([4a124b7](https://github.com/VU-ASE/rover/commit/4a124b7ddec97a31e733b064745a4374e4cd71fc))
* start roverctl-web container ([9329015](https://github.com/VU-ASE/rover/commit/9329015b9cb86a8db86bc756052327fdf170785e))
* starting, stopping and modifying pipelines ([e54a176](https://github.com/VU-ASE/rover/commit/e54a1762fdc686ec0236874c1169ea9c20014bff))
* unified toasts ([ded557b](https://github.com/VU-ASE/rover/commit/ded557be905b07b688a2fe9f6098a028543ec8bb))
* update endpoints ([e1422b4](https://github.com/VU-ASE/rover/commit/e1422b40ed2e25e98eb094bc088154d896a28caf))
* update target ([0f3cc2a](https://github.com/VU-ASE/rover/commit/0f3cc2a340b5a78a3c04fc363941608fa9622613))
* upgrade roverctl to go 1.22.12 ([6a19635](https://github.com/VU-ASE/rover/commit/6a1963502affb37a09328151d1c55622ffb14d0f))
* use mDNS to resolve hostnames, allow proxy IP override ([2cb436e](https://github.com/VU-ASE/rover/commit/2cb436e2038e2ef1e5d722913f0d794ce9d60bfe))
* warn users if roverctl mismatches roverd and install appr roverctl ([868e3c1](https://github.com/VU-ASE/rover/commit/868e3c1db16c7dc0955121ae022ded73ff152f0a))


### Bug Fixes

* added service-template for cpp ([cf5fd0c](https://github.com/VU-ASE/rover/commit/cf5fd0c04b646dfef25138bbbb54220c03a2722e))
* added smoother steering to testing scripts ([681d881](https://github.com/VU-ASE/rover/commit/681d881c2147039cd8d2954098833fd0110bc080))
* auto-update roverd in test ([9b53863](https://github.com/VU-ASE/rover/commit/9b53863f55a2baf712b25f46099b1d2bbd7b7ff3))
* broken links ([7a4d2aa](https://github.com/VU-ASE/rover/commit/7a4d2aa8ee9bc87d5f06ea3c1f55a9f324cc57a7))
* build to ./bin ([dba0f2b](https://github.com/VU-ASE/rover/commit/dba0f2b1cfd4deaa257beb7c0bbe241b02879f5d))
* ci testing ([269b7a3](https://github.com/VU-ASE/rover/commit/269b7a3b9ece651b79b22145db9eccb1cf018fa9))
* consistent style for services ([a9f5fbe](https://github.com/VU-ASE/rover/commit/a9f5fbef311b50ce30ca3891a9c5591ec3c12d91))
* correct link if upload fails ([d229edb](https://github.com/VU-ASE/rover/commit/d229edbeca49c5684dd881e97dc9d02bf9cbb9fc))
* correct timing of service error messages to actual moment of crash ([68f158f](https://github.com/VU-ASE/rover/commit/68f158f9404125399fc18499d855f04f6ab054f0))
* create openapi directory if it does not exist ([1f881b8](https://github.com/VU-ASE/rover/commit/1f881b8846f38ceea597ab31be031ae7c60ecf96))
* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* disable transceiver service when debug mode is not enabled ([46f41e5](https://github.com/VU-ASE/rover/commit/46f41e5f75025b3dd3515852eb42f69f88afccea))
* display warning if author is not set in info command ([b2c5eaa](https://github.com/VU-ASE/rover/commit/b2c5eaa5e866c87066cdd2eddc1aecd6f680d493))
* do not build roverd on self-hosted ([af62651](https://github.com/VU-ASE/rover/commit/af626516335e3f1c2a5682aaa8fb54b5f72deb5c))
* do not error on existing connection being reused ([6e357bf](https://github.com/VU-ASE/rover/commit/6e357bfd6dc7c008d09f0717992a3619f7f2240c))
* do not publish openapi files ([f7f7ff0](https://github.com/VU-ASE/rover/commit/f7f7ff039253c5ebfc893fa0d53b5daad304968d))
* do not reset calibration value for normal pipeline testing ([be21d46](https://github.com/VU-ASE/rover/commit/be21d462d992f50a90e93b473417e6a0f026b3f3))
* do not show "shortcuts" tab as active on debugpage ([822b084](https://github.com/VU-ASE/rover/commit/822b084f1a15a00c34a12a162d8c618e61bc5c65))
* do NOT update roverd in the test script ([b5f1cb0](https://github.com/VU-ASE/rover/commit/b5f1cb0ea34f527117f6aee6a8d071002786c52d))
* docs and install script for roverctl ([a069205](https://github.com/VU-ASE/rover/commit/a069205b726187a5ad7ad99c8f00a75f7569ce9f))
* emergency stop ([2b7bdf8](https://github.com/VU-ASE/rover/commit/2b7bdf85012d4b5e736a5bf3543a7f5856bc9855))
* error on roverctl install.sh if release does not exist or no binary ([1880064](https://github.com/VU-ASE/rover/commit/1880064e54514a502ae07bd21f690b8d953a7eea))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* extract artifact downloads first ([37762b7](https://github.com/VU-ASE/rover/commit/37762b77e75becc3a8fb960e83ddd72f47305729))
* extract from correct location ([d4106ba](https://github.com/VU-ASE/rover/commit/d4106ba17d7e8ec4d7fea3f3827aded9d4a25d0a))
* gitignore was missing readme ([44c64ff](https://github.com/VU-ASE/rover/commit/44c64ffed3c5d317704599bfb95af5ace86af280))
* hard-coded version ([058573d](https://github.com/VU-ASE/rover/commit/058573d9cf67a814ef554712b4ad610eaabcb0b4))
* hide updater screen when all services are installed ([f4dce16](https://github.com/VU-ASE/rover/commit/f4dce16e2ae91666336fa361d6c2d15db813c380))
* image push only succeeds on --verbose flag, port checking ([b0d5998](https://github.com/VU-ASE/rover/commit/b0d599895c6cc1f3365eca1aee7ceadbadb6e13e))
* improved upload description and example usage ([248002c](https://github.com/VU-ASE/rover/commit/248002c3d1c96494a91577621a270ea0a3a68cea))
* incorrect number of arguments for service init script ([49981a0](https://github.com/VU-ASE/rover/commit/49981a094084063c1f584caf0f3df2a25e13aeca))
* infinite redirect ([0f63b01](https://github.com/VU-ASE/rover/commit/0f63b019402c498c496f4cec1bbf9f483de3065a))
* load list of services from roverd in debug window ([a8099bb](https://github.com/VU-ASE/rover/commit/a8099bb929543b3dfde14b06bc61115c880f3ceb))
* make roverctl compatible with new API spec ([6eb675c](https://github.com/VU-ASE/rover/commit/6eb675c87be7b49d31ca3405906767df359405f3))
* merge service name and service "as" name in tuning overview ([3675c0e](https://github.com/VU-ASE/rover/commit/3675c0e5159f0e503c8ff65a596c5f757263a9c5))
* more helpful messages to enable debug mode ([106e3ea](https://github.com/VU-ASE/rover/commit/106e3ea929bc7943948e77f07074171e4a31ad55))
* more lenient timeouts for roverd downloads ([a89b302](https://github.com/VU-ASE/rover/commit/a89b302fd97b48ce176d1f27becee737a67ac803))
* move github.com/VU-ASE/roverctl to github.com/VU-ASE/rover/roverctl ([a35e0dc](https://github.com/VU-ASE/rover/commit/a35e0dc12895750545256432284054e560e9eb65))
* navbar type error ([d66c700](https://github.com/VU-ASE/rover/commit/d66c70009adc0a222eba5967af53d168de2f44af))
* only show services when the pipeline is started (in debug mode) ([e22facf](https://github.com/VU-ASE/rover/commit/e22facf86deb8c7b305238dcc704b6aab7fc741b))
* open-api-gen debugging ([013bce9](https://github.com/VU-ASE/rover/commit/013bce9f4a0838ccf91ddb53e429d27cb7132178))
* permission issues for building ([8f82d02](https://github.com/VU-ASE/rover/commit/8f82d02b4b4bcd6a5930c5faa8fd504042a69f0b))
* permissions on github actions ([0b7225a](https://github.com/VU-ASE/rover/commit/0b7225ae596dd3d1e1a29a46ad26b2f6e256d9dd))
* pipeline not loading after viewing debug screen ([67787c4](https://github.com/VU-ASE/rover/commit/67787c480eb8d8c92381433008e21b2297b5b6eb))
* point install scripts to rover repository (was roverctl) ([b139324](https://github.com/VU-ASE/rover/commit/b139324d303d27c88f2f8d91d0aaf1b78a5f714a))
* poll pipeline status ([a02bbf6](https://github.com/VU-ASE/rover/commit/a02bbf698dbd8888b4f8efcf82dfb2fdb7eca9a8))
* proxy server deadlock on repeated transceiver crashes ([577a1a2](https://github.com/VU-ASE/rover/commit/577a1a29775debc7e75526abbc9f799e9b3db0ec))
* public link for schematic in docs ([b26decf](https://github.com/VU-ASE/rover/commit/b26decf57eb08b15e0fd4c2c04e81d1f9b3a1fdd))
* release assets on all success ([4448019](https://github.com/VU-ASE/rover/commit/44480199fcac3cc9c81e317ef7b34c270641ec6a))
* remove unused github action jobs ([04ef005](https://github.com/VU-ASE/rover/commit/04ef0051fa635a52e6fbc521c6ca915b58c1423e))
* remove unused view ([31f91cd](https://github.com/VU-ASE/rover/commit/31f91cd1bd482a2a1bb05af107ebb19b153d1de5))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* rename passthrough to appropriate folder ([5c1d41e](https://github.com/VU-ASE/rover/commit/5c1d41e8f3a1bb3be425841ede4c9a80fe2c2861))
* retain debug toggle information and don't show unusable tuning ([9d78947](https://github.com/VU-ASE/rover/commit/9d789474196bc613cea17b1da0a775e8669e9d12))
* returning service_as optimistically for all endpoints ([a019fe1](https://github.com/VU-ASE/rover/commit/a019fe1bee5a0a5e591911b389ee599b647e0440))
* roverctl --force warning message ([5384c70](https://github.com/VU-ASE/rover/commit/5384c704cac253524bc7d433d4b260602e1c3632))
* roverctl install.sh, verify downloaded binary is valid ([381a5ed](https://github.com/VU-ASE/rover/commit/381a5ed2e809710a882a65e4affd3fb658ecb2e2))
* roverctl replace FQN types ([f6a9c44](https://github.com/VU-ASE/rover/commit/f6a9c44dee576b4ba63f3e44ca5a07704ead245b))
* roverctl-roverd update API endpoint not working ([9fd776f](https://github.com/VU-ASE/rover/commit/9fd776f4616d2c66e59a66cb0c2b554620a72522))
* roverctl, build openapi files in Dockerfile ([af01eb6](https://github.com/VU-ASE/rover/commit/af01eb620b565518860cb84d5665595788b0de71))
* roverd early check for root ([451f7eb](https://github.com/VU-ASE/rover/commit/451f7eb4ab031fe82dc0d7b19e2a7377122a2e5f))
* roverd log lines ([97adc00](https://github.com/VU-ASE/rover/commit/97adc0044f551ce2aa9ea1e7db9beeae24e43ab7))
* roverd not propagating stream outputs to transceiver if no consumer ([0818f38](https://github.com/VU-ASE/rover/commit/0818f38c0b4353bb5078a3c8b904be8a9e5f3c38))
* roverd shows all exit codes now ([2de0226](https://github.com/VU-ASE/rover/commit/2de022691173795eac6a63eec9787d554f466958))
* roverd shutsdown pipeline on termination ([83ea1dc](https://github.com/VU-ASE/rover/commit/83ea1dce347970d9a265472b62a18031177b46eb))
* roverd unclosed delimiter ([17d8e95](https://github.com/VU-ASE/rover/commit/17d8e95a9b322a205ffe46c5dd367459f59c9aef))
* run open-api generator as devuser ([e8e84c8](https://github.com/VU-ASE/rover/commit/e8e84c85578d578299798f469782f04e9acff0f0))
* service logs cut-off point makes them unreadable ([412a8ee](https://github.com/VU-ASE/rover/commit/412a8ee5eb7520daee8fbef481b0271eadab9470))
* service yamls can now specify a bash command for run ([6022928](https://github.com/VU-ASE/rover/commit/60229285d294b95698faef9958704e3d8c81ff5e))
* show correct fix for updating roverd on mismatch ([a49144f](https://github.com/VU-ASE/rover/commit/a49144feec173a050d84a6bfee482bcb7fe7afc3))
* show correct panels for selected services ([5ba569e](https://github.com/VU-ASE/rover/commit/5ba569e06f205686fd5baa2dd7ff0ae765d06bb7))
* show debug toggle correctly ([3ea63f8](https://github.com/VU-ASE/rover/commit/3ea63f8a79ee4fe7ca0836407fb96d6f501fd94d))
* show logs of last run when selected ([0cca2da](https://github.com/VU-ASE/rover/commit/0cca2dab6e762f6fb8916ed3cce3175f17262ec0))
* show pipeline edges accurately ([ff53e34](https://github.com/VU-ASE/rover/commit/ff53e3424e6c8644bd943c3d181c080cf69f6a74))
* show service crash exit code on transition from started to stopped ([e43cb26](https://github.com/VU-ASE/rover/commit/e43cb260ef1df98869502afda7731fe188257beb))
* sort services alphabetically. Always show vu-ase first ([fb6a58b](https://github.com/VU-ASE/rover/commit/fb6a58b9419af46bb1b06792e548bdc9431fb6da))
* specifying value types in service.yaml is required ([e8757a3](https://github.com/VU-ASE/rover/commit/e8757a3edeb67edcccb0cd5d6909d15879d6f886))
* testing ci again ([573bb0f](https://github.com/VU-ASE/rover/commit/573bb0f5aafb6ca7b18208a4c71b1a14724de484))
* transceiver output address should use `*` instead of `localhost` ([19264da](https://github.com/VU-ASE/rover/commit/19264dac0d2c18e4ff8485889e4f9dd0dd0fddab))
* update checking for default pipeline ([3f67220](https://github.com/VU-ASE/rover/commit/3f672204d7faaeae6ec4a3d2ec141ab52af4dc55))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))
* use correct version of action/artifact ([300beda](https://github.com/VU-ASE/rover/commit/300beda2031241877d4101279e5f9ec5cf8dcd07))
* use docker buildx for builds ([aa4893b](https://github.com/VU-ASE/rover/commit/aa4893b9e97492b82198164711f7cca2e534f094))
* use dynamic env variables for configuration ([25d5c88](https://github.com/VU-ASE/rover/commit/25d5c88b217c9040aea12686a4b08b05443a8e03))
* use lowercase for GHCR tag ([e62a583](https://github.com/VU-ASE/rover/commit/e62a58302e17c7ca9739349b60fc654cce7f7da3))
* use mdns for tests ([bb59b70](https://github.com/VU-ASE/rover/commit/bb59b70cbbbc7e8a31fdf4e93a85f63346c22e5f))
* using mv instead of os.Rename ([696ecc4](https://github.com/VU-ASE/rover/commit/696ecc440505fb39c3ec5cf814b55786b9fb5bf5))
* version mismatch checking ([4ba63b9](https://github.com/VU-ASE/rover/commit/4ba63b901f88055e68db2a7234ca17659588798c))
* version update script now accepts "v" prefix ([362c24c](https://github.com/VU-ASE/rover/commit/362c24c0b3264ce7d3aef7fdfabbfb6bea72d45a))
* version utility and correct comparison ([9251512](https://github.com/VU-ASE/rover/commit/92515124a450c2edd2ad8aed10a20b7b42ba4b9c))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))
* release 0.0.1 ([abd806e](https://github.com/VU-ASE/rover/commit/abd806e59353ee5ea9da82a80612ff9ece653285))
* release 0.0.5 ([aa2ed11](https://github.com/VU-ASE/rover/commit/aa2ed11e18453f8a6edc3f82cd503e38288e7060))
* release 1.0.0 ([55d09fd](https://github.com/VU-ASE/rover/commit/55d09fd29d0bc18c528ce63af45a2c10c3428f8d))

## [1.0.0](https://github.com/VU-ASE/rover/compare/v1.0.0...v1.0.0) (2025-04-01)


### ⚠ BREAKING CHANGES

* release rover v1.0.0

### Features

* add alias "rover" to roverctl install script ([1ec8389](https://github.com/VU-ASE/rover/commit/1ec83893a0a7458aef2c8b41791c548ffda961df))
* add incompatibility warning to all endpoints that need it ([126849c](https://github.com/VU-ASE/rover/commit/126849cd89e743d089f2765f7261df24c30fd0eb))
* add installation location to service info view ([1a0ca7c](https://github.com/VU-ASE/rover/commit/1a0ca7c58c0a9b58ef3a0073babc92470ce99fe1))
* add passthrough back ([e07dc70](https://github.com/VU-ASE/rover/commit/e07dc7073c6b63dd9ed1c9894a11df2127e7d23c))
* add shortcuts item in navbar ([5aea4d3](https://github.com/VU-ASE/rover/commit/5aea4d3c7ec4458cd4aa6feb79e62398d5c1d860))
* added /emergency endpoint ([55df2dc](https://github.com/VU-ASE/rover/commit/55df2dcf7d51e36a6bec730cd36b2324de9cd45b))
* added Dockerfile for roverctl-web ([ec9733b](https://github.com/VU-ASE/rover/commit/ec9733b2f507d03d0ec405e6bcc2d82519406e6b))
* added manage script ([e1b1429](https://github.com/VU-ASE/rover/commit/e1b142989b1b0c0dfcaa55f0a496da56031288e4))
* added passthrough and debugging functionality ([90975ef](https://github.com/VU-ASE/rover/commit/90975ef34e1acf2ffc1f9ec740e18b600df7d561))
* api adaptation ([d36262f](https://github.com/VU-ASE/rover/commit/d36262f0e0d336e08cc1a3a6e2484ad4e52c2431))
* api clean up and addition of alias error (BBC) ([1d685a6](https://github.com/VU-ASE/rover/commit/1d685a6578ddc0b9c3b48fa9ffeb0abecfefad22))
* app config and overview page ([1c76d8c](https://github.com/VU-ASE/rover/commit/1c76d8cd42abe86a77e5133806ac97921a4da5d5))
* author command ([6ead41c](https://github.com/VU-ASE/rover/commit/6ead41cdbfb58ab0ce484cb627ae01bae0223da1))
* build command ([ea610eb](https://github.com/VU-ASE/rover/commit/ea610eb0171be57fbfd41ca2991bc6a1d00dedea))
* calibration command for actuator service ([a549e94](https://github.com/VU-ASE/rover/commit/a549e9499f9b629e89b40ea6db1fc4e1e6f025c4))
* CLI approach to roverctl ([bb78304](https://github.com/VU-ASE/rover/commit/bb78304cf2c6fcda4b2ce95bba51fcd183ee6ce3))
* dangers page ([a296ad3](https://github.com/VU-ASE/rover/commit/a296ad3ccb2d682d7a881c057d5a60aec36a9888))
* default ASE pipeline updates and installs ([0abcdb0](https://github.com/VU-ASE/rover/commit/0abcdb0021d0c2c4158a52e479684615864518af))
* dialog and upgrad/downgrade roverctl on version mismatch ([5d25e52](https://github.com/VU-ASE/rover/commit/5d25e52df3a75dfad549aca7e929f39d0755aa6f))
* emergency command ([2b883cd](https://github.com/VU-ASE/rover/commit/2b883cd7ded78acf211950161ec854884722badd))
* enable and install debug logic from roverctl-web ([a760e30](https://github.com/VU-ASE/rover/commit/a760e30d6991e0322399aeb4d65beca78a06dcc6))
* error types for the configuration endpoint ([016a94b](https://github.com/VU-ASE/rover/commit/016a94bd55a8643d300bf94d1c88c54bdb60512b))
* github actions workflows to build and push docker images ([5c4b035](https://github.com/VU-ASE/rover/commit/5c4b035329c9998331228113913a28fcc72b7552))
* group test scripts ([d073bfd](https://github.com/VU-ASE/rover/commit/d073bfdecdd198ae74f7b697144ed7dee806826c))
* helpful toasts and warning messages when pipeline starts/stops ([27074f6](https://github.com/VU-ASE/rover/commit/27074f67ca459497ca9105b6ec2c419f70a2d49b))
* implemented in place update API ([44a8c73](https://github.com/VU-ASE/rover/commit/44a8c73ffd2ad93d22056335e78b3a8ff4a5da39))
* improved error printing with multiline support ([8c64c2c](https://github.com/VU-ASE/rover/commit/8c64c2ce5443774ab1571658d5f57e932cb426b1))
* improved error reporting on upload failure ([b657f17](https://github.com/VU-ASE/rover/commit/b657f17d60bdaf464a2e998ddbb1541fead96141))
* info command ([478704e](https://github.com/VU-ASE/rover/commit/478704ec6fa6d5998f16da28bf9116df0c3dac45))
* info target ([0bc9cb0](https://github.com/VU-ASE/rover/commit/0bc9cb0ecfd0d6ffe2aac31c12717839d7ede73f))
* install command ([955fea9](https://github.com/VU-ASE/rover/commit/955fea970534b51d0e8e31abef62600f776475c1))
* install script that allows for installing a specific version ([f5bfcaf](https://github.com/VU-ASE/rover/commit/f5bfcaf2bada29e9e8732b61d79c93a53c3331d9))
* install script with symlink instead of alias ([22ccd67](https://github.com/VU-ASE/rover/commit/22ccd6786c112cbde2979064db8ae6ef89fd5c46))
* keyboard shortcuts for emergency, stopping pipelin and debug pause ([ad50f10](https://github.com/VU-ASE/rover/commit/ad50f1014705f1631c973928ff06922c74733514))
* link source code and docs for official services ([3ac9131](https://github.com/VU-ASE/rover/commit/3ac9131e4fce750e5475ec18167cf9feb83109c1))
* logs endpoint ([f1b3077](https://github.com/VU-ASE/rover/commit/f1b3077fb46bcc396dc9a7c1d32c96b8a66537d7))
* manage page (wip) and status overlay ([e9132ed](https://github.com/VU-ASE/rover/commit/e9132ed41535b49fbc10353c023d6935edf82640))
* merge passthrough into roverctl, debugging now works ([8f6c88f](https://github.com/VU-ASE/rover/commit/8f6c88f43323856971c7f6f129c47a0d23ba3a1e))
* pin version dependency between roverctl-roverd and warn if broken ([c8e98dc](https://github.com/VU-ASE/rover/commit/c8e98dcc4b89d079625197e91685f5b085be4bfd))
* pipeline enable/disable endpoints ([148fb7e](https://github.com/VU-ASE/rover/commit/148fb7eb16da42ae3af8ce568a767419260d9704))
* pipeline endpoint ([d0beca5](https://github.com/VU-ASE/rover/commit/d0beca50b2d01d0cedaada7833edd6a8bcba1053))
* release rover v1.0.0 ([292c229](https://github.com/VU-ASE/rover/commit/292c22939aeb4b8592c2dbe8ac1566c350d83b1a))
* reset calibration values on trimming ([ec8eb15](https://github.com/VU-ASE/rover/commit/ec8eb15b7efc09a2b245a1424f21fb4bb51beb9d))
* roverctl service init endpoint ([b8a9027](https://github.com/VU-ASE/rover/commit/b8a90274bc33f681fc5eaf97cb8adf04b828122f))
* roverctl-web startup through docker ([e3ca9a5](https://github.com/VU-ASE/rover/commit/e3ca9a54be65233a5f7a8f8fa8b915c9d042b724))
* service delete endpoint ([185362f](https://github.com/VU-ASE/rover/commit/185362f0f91d3af1cac2b0abe58e4741260740f9))
* service information ([c0119a5](https://github.com/VU-ASE/rover/commit/c0119a5c4715cb39d86fdecdf856ce3d48127595))
* service installation ([246a452](https://github.com/VU-ASE/rover/commit/246a452e0ae9b4d8489b1f852d1a425c7137e895))
* service selection toggles and ordening ([6ad02f4](https://github.com/VU-ASE/rover/commit/6ad02f49222b35caec422dc1e2c486276a50aa65))
* service_as feature ([782ab78](https://github.com/VU-ASE/rover/commit/782ab7874c1eb6eec0e58ac73283f2f9ab940a5b))
* services endpoint ([a5e11a4](https://github.com/VU-ASE/rover/commit/a5e11a46ca7d8c5a20345a356f32943e86ba50f0))
* shorthand aliasses for most commands ([3c270ce](https://github.com/VU-ASE/rover/commit/3c270ce7e0fbff41f1273a2cedd927f88d24fb99))
* show transceiver in roverctl-web ([be2cb92](https://github.com/VU-ASE/rover/commit/be2cb9274d7bed0c2d4e890b33018c8f82fbd652))
* show which service crashed the pipeline ([8591295](https://github.com/VU-ASE/rover/commit/85912954080cdb6654596702b195fb1c2b064ded))
* shutdown command ([4a124b7](https://github.com/VU-ASE/rover/commit/4a124b7ddec97a31e733b064745a4374e4cd71fc))
* start roverctl-web container ([9329015](https://github.com/VU-ASE/rover/commit/9329015b9cb86a8db86bc756052327fdf170785e))
* starting, stopping and modifying pipelines ([e54a176](https://github.com/VU-ASE/rover/commit/e54a1762fdc686ec0236874c1169ea9c20014bff))
* unified toasts ([ded557b](https://github.com/VU-ASE/rover/commit/ded557be905b07b688a2fe9f6098a028543ec8bb))
* update endpoints ([e1422b4](https://github.com/VU-ASE/rover/commit/e1422b40ed2e25e98eb094bc088154d896a28caf))
* update target ([0f3cc2a](https://github.com/VU-ASE/rover/commit/0f3cc2a340b5a78a3c04fc363941608fa9622613))
* upgrade roverctl to go 1.22.12 ([6a19635](https://github.com/VU-ASE/rover/commit/6a1963502affb37a09328151d1c55622ffb14d0f))
* use mDNS to resolve hostnames, allow proxy IP override ([2cb436e](https://github.com/VU-ASE/rover/commit/2cb436e2038e2ef1e5d722913f0d794ce9d60bfe))
* warn users if roverctl mismatches roverd and install appr roverctl ([868e3c1](https://github.com/VU-ASE/rover/commit/868e3c1db16c7dc0955121ae022ded73ff152f0a))


### Bug Fixes

* added service-template for cpp ([cf5fd0c](https://github.com/VU-ASE/rover/commit/cf5fd0c04b646dfef25138bbbb54220c03a2722e))
* added smoother steering to testing scripts ([681d881](https://github.com/VU-ASE/rover/commit/681d881c2147039cd8d2954098833fd0110bc080))
* auto-update roverd in test ([9b53863](https://github.com/VU-ASE/rover/commit/9b53863f55a2baf712b25f46099b1d2bbd7b7ff3))
* broken links ([7a4d2aa](https://github.com/VU-ASE/rover/commit/7a4d2aa8ee9bc87d5f06ea3c1f55a9f324cc57a7))
* build to ./bin ([dba0f2b](https://github.com/VU-ASE/rover/commit/dba0f2b1cfd4deaa257beb7c0bbe241b02879f5d))
* ci testing ([269b7a3](https://github.com/VU-ASE/rover/commit/269b7a3b9ece651b79b22145db9eccb1cf018fa9))
* consistent style for services ([a9f5fbe](https://github.com/VU-ASE/rover/commit/a9f5fbef311b50ce30ca3891a9c5591ec3c12d91))
* correct link if upload fails ([d229edb](https://github.com/VU-ASE/rover/commit/d229edbeca49c5684dd881e97dc9d02bf9cbb9fc))
* correct timing of service error messages to actual moment of crash ([68f158f](https://github.com/VU-ASE/rover/commit/68f158f9404125399fc18499d855f04f6ab054f0))
* create openapi directory if it does not exist ([1f881b8](https://github.com/VU-ASE/rover/commit/1f881b8846f38ceea597ab31be031ae7c60ecf96))
* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* disable transceiver service when debug mode is not enabled ([46f41e5](https://github.com/VU-ASE/rover/commit/46f41e5f75025b3dd3515852eb42f69f88afccea))
* display warning if author is not set in info command ([b2c5eaa](https://github.com/VU-ASE/rover/commit/b2c5eaa5e866c87066cdd2eddc1aecd6f680d493))
* do not build roverd on self-hosted ([af62651](https://github.com/VU-ASE/rover/commit/af626516335e3f1c2a5682aaa8fb54b5f72deb5c))
* do not error on existing connection being reused ([6e357bf](https://github.com/VU-ASE/rover/commit/6e357bfd6dc7c008d09f0717992a3619f7f2240c))
* do not publish openapi files ([f7f7ff0](https://github.com/VU-ASE/rover/commit/f7f7ff039253c5ebfc893fa0d53b5daad304968d))
* do not reset calibration value for normal pipeline testing ([be21d46](https://github.com/VU-ASE/rover/commit/be21d462d992f50a90e93b473417e6a0f026b3f3))
* do not show "shortcuts" tab as active on debugpage ([822b084](https://github.com/VU-ASE/rover/commit/822b084f1a15a00c34a12a162d8c618e61bc5c65))
* do NOT update roverd in the test script ([b5f1cb0](https://github.com/VU-ASE/rover/commit/b5f1cb0ea34f527117f6aee6a8d071002786c52d))
* docs and install script for roverctl ([a069205](https://github.com/VU-ASE/rover/commit/a069205b726187a5ad7ad99c8f00a75f7569ce9f))
* emergency stop ([2b7bdf8](https://github.com/VU-ASE/rover/commit/2b7bdf85012d4b5e736a5bf3543a7f5856bc9855))
* error on roverctl install.sh if release does not exist or no binary ([1880064](https://github.com/VU-ASE/rover/commit/1880064e54514a502ae07bd21f690b8d953a7eea))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* extract artifact downloads first ([37762b7](https://github.com/VU-ASE/rover/commit/37762b77e75becc3a8fb960e83ddd72f47305729))
* extract from correct location ([d4106ba](https://github.com/VU-ASE/rover/commit/d4106ba17d7e8ec4d7fea3f3827aded9d4a25d0a))
* gitignore was missing readme ([44c64ff](https://github.com/VU-ASE/rover/commit/44c64ffed3c5d317704599bfb95af5ace86af280))
* hard-coded version ([058573d](https://github.com/VU-ASE/rover/commit/058573d9cf67a814ef554712b4ad610eaabcb0b4))
* hide updater screen when all services are installed ([f4dce16](https://github.com/VU-ASE/rover/commit/f4dce16e2ae91666336fa361d6c2d15db813c380))
* image push only succeeds on --verbose flag, port checking ([b0d5998](https://github.com/VU-ASE/rover/commit/b0d599895c6cc1f3365eca1aee7ceadbadb6e13e))
* improved upload description and example usage ([248002c](https://github.com/VU-ASE/rover/commit/248002c3d1c96494a91577621a270ea0a3a68cea))
* incorrect number of arguments for service init script ([49981a0](https://github.com/VU-ASE/rover/commit/49981a094084063c1f584caf0f3df2a25e13aeca))
* infinite redirect ([0f63b01](https://github.com/VU-ASE/rover/commit/0f63b019402c498c496f4cec1bbf9f483de3065a))
* load list of services from roverd in debug window ([a8099bb](https://github.com/VU-ASE/rover/commit/a8099bb929543b3dfde14b06bc61115c880f3ceb))
* make roverctl compatible with new API spec ([6eb675c](https://github.com/VU-ASE/rover/commit/6eb675c87be7b49d31ca3405906767df359405f3))
* merge service name and service "as" name in tuning overview ([3675c0e](https://github.com/VU-ASE/rover/commit/3675c0e5159f0e503c8ff65a596c5f757263a9c5))
* more helpful messages to enable debug mode ([106e3ea](https://github.com/VU-ASE/rover/commit/106e3ea929bc7943948e77f07074171e4a31ad55))
* more lenient timeouts for roverd downloads ([a89b302](https://github.com/VU-ASE/rover/commit/a89b302fd97b48ce176d1f27becee737a67ac803))
* move github.com/VU-ASE/roverctl to github.com/VU-ASE/rover/roverctl ([a35e0dc](https://github.com/VU-ASE/rover/commit/a35e0dc12895750545256432284054e560e9eb65))
* navbar type error ([d66c700](https://github.com/VU-ASE/rover/commit/d66c70009adc0a222eba5967af53d168de2f44af))
* only show services when the pipeline is started (in debug mode) ([e22facf](https://github.com/VU-ASE/rover/commit/e22facf86deb8c7b305238dcc704b6aab7fc741b))
* open-api-gen debugging ([013bce9](https://github.com/VU-ASE/rover/commit/013bce9f4a0838ccf91ddb53e429d27cb7132178))
* permission issues for building ([8f82d02](https://github.com/VU-ASE/rover/commit/8f82d02b4b4bcd6a5930c5faa8fd504042a69f0b))
* permissions on github actions ([0b7225a](https://github.com/VU-ASE/rover/commit/0b7225ae596dd3d1e1a29a46ad26b2f6e256d9dd))
* pipeline not loading after viewing debug screen ([67787c4](https://github.com/VU-ASE/rover/commit/67787c480eb8d8c92381433008e21b2297b5b6eb))
* point install scripts to rover repository (was roverctl) ([b139324](https://github.com/VU-ASE/rover/commit/b139324d303d27c88f2f8d91d0aaf1b78a5f714a))
* poll pipeline status ([a02bbf6](https://github.com/VU-ASE/rover/commit/a02bbf698dbd8888b4f8efcf82dfb2fdb7eca9a8))
* proxy server deadlock on repeated transceiver crashes ([577a1a2](https://github.com/VU-ASE/rover/commit/577a1a29775debc7e75526abbc9f799e9b3db0ec))
* public link for schematic in docs ([b26decf](https://github.com/VU-ASE/rover/commit/b26decf57eb08b15e0fd4c2c04e81d1f9b3a1fdd))
* release assets on all success ([4448019](https://github.com/VU-ASE/rover/commit/44480199fcac3cc9c81e317ef7b34c270641ec6a))
* remove unused github action jobs ([04ef005](https://github.com/VU-ASE/rover/commit/04ef0051fa635a52e6fbc521c6ca915b58c1423e))
* remove unused view ([31f91cd](https://github.com/VU-ASE/rover/commit/31f91cd1bd482a2a1bb05af107ebb19b153d1de5))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* rename passthrough to appropriate folder ([5c1d41e](https://github.com/VU-ASE/rover/commit/5c1d41e8f3a1bb3be425841ede4c9a80fe2c2861))
* retain debug toggle information and don't show unusable tuning ([9d78947](https://github.com/VU-ASE/rover/commit/9d789474196bc613cea17b1da0a775e8669e9d12))
* returning service_as optimistically for all endpoints ([a019fe1](https://github.com/VU-ASE/rover/commit/a019fe1bee5a0a5e591911b389ee599b647e0440))
* roverctl --force warning message ([5384c70](https://github.com/VU-ASE/rover/commit/5384c704cac253524bc7d433d4b260602e1c3632))
* roverctl install.sh, verify downloaded binary is valid ([381a5ed](https://github.com/VU-ASE/rover/commit/381a5ed2e809710a882a65e4affd3fb658ecb2e2))
* roverctl replace FQN types ([f6a9c44](https://github.com/VU-ASE/rover/commit/f6a9c44dee576b4ba63f3e44ca5a07704ead245b))
* roverctl-roverd update API endpoint not working ([9fd776f](https://github.com/VU-ASE/rover/commit/9fd776f4616d2c66e59a66cb0c2b554620a72522))
* roverctl, build openapi files in Dockerfile ([af01eb6](https://github.com/VU-ASE/rover/commit/af01eb620b565518860cb84d5665595788b0de71))
* roverd early check for root ([451f7eb](https://github.com/VU-ASE/rover/commit/451f7eb4ab031fe82dc0d7b19e2a7377122a2e5f))
* roverd log lines ([97adc00](https://github.com/VU-ASE/rover/commit/97adc0044f551ce2aa9ea1e7db9beeae24e43ab7))
* roverd not propagating stream outputs to transceiver if no consumer ([0818f38](https://github.com/VU-ASE/rover/commit/0818f38c0b4353bb5078a3c8b904be8a9e5f3c38))
* roverd shows all exit codes now ([2de0226](https://github.com/VU-ASE/rover/commit/2de022691173795eac6a63eec9787d554f466958))
* roverd shutsdown pipeline on termination ([83ea1dc](https://github.com/VU-ASE/rover/commit/83ea1dce347970d9a265472b62a18031177b46eb))
* roverd unclosed delimiter ([17d8e95](https://github.com/VU-ASE/rover/commit/17d8e95a9b322a205ffe46c5dd367459f59c9aef))
* run open-api generator as devuser ([e8e84c8](https://github.com/VU-ASE/rover/commit/e8e84c85578d578299798f469782f04e9acff0f0))
* service logs cut-off point makes them unreadable ([412a8ee](https://github.com/VU-ASE/rover/commit/412a8ee5eb7520daee8fbef481b0271eadab9470))
* service yamls can now specify a bash command for run ([6022928](https://github.com/VU-ASE/rover/commit/60229285d294b95698faef9958704e3d8c81ff5e))
* show correct fix for updating roverd on mismatch ([a49144f](https://github.com/VU-ASE/rover/commit/a49144feec173a050d84a6bfee482bcb7fe7afc3))
* show correct panels for selected services ([5ba569e](https://github.com/VU-ASE/rover/commit/5ba569e06f205686fd5baa2dd7ff0ae765d06bb7))
* show debug toggle correctly ([3ea63f8](https://github.com/VU-ASE/rover/commit/3ea63f8a79ee4fe7ca0836407fb96d6f501fd94d))
* show logs of last run when selected ([0cca2da](https://github.com/VU-ASE/rover/commit/0cca2dab6e762f6fb8916ed3cce3175f17262ec0))
* show pipeline edges accurately ([ff53e34](https://github.com/VU-ASE/rover/commit/ff53e3424e6c8644bd943c3d181c080cf69f6a74))
* show service crash exit code on transition from started to stopped ([e43cb26](https://github.com/VU-ASE/rover/commit/e43cb260ef1df98869502afda7731fe188257beb))
* sort services alphabetically. Always show vu-ase first ([fb6a58b](https://github.com/VU-ASE/rover/commit/fb6a58b9419af46bb1b06792e548bdc9431fb6da))
* specifying value types in service.yaml is required ([e8757a3](https://github.com/VU-ASE/rover/commit/e8757a3edeb67edcccb0cd5d6909d15879d6f886))
* testing ci again ([573bb0f](https://github.com/VU-ASE/rover/commit/573bb0f5aafb6ca7b18208a4c71b1a14724de484))
* transceiver output address should use `*` instead of `localhost` ([19264da](https://github.com/VU-ASE/rover/commit/19264dac0d2c18e4ff8485889e4f9dd0dd0fddab))
* update checking for default pipeline ([3f67220](https://github.com/VU-ASE/rover/commit/3f672204d7faaeae6ec4a3d2ec141ab52af4dc55))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))
* use correct version of action/artifact ([300beda](https://github.com/VU-ASE/rover/commit/300beda2031241877d4101279e5f9ec5cf8dcd07))
* use docker buildx for builds ([aa4893b](https://github.com/VU-ASE/rover/commit/aa4893b9e97492b82198164711f7cca2e534f094))
* use dynamic env variables for configuration ([25d5c88](https://github.com/VU-ASE/rover/commit/25d5c88b217c9040aea12686a4b08b05443a8e03))
* use lowercase for GHCR tag ([e62a583](https://github.com/VU-ASE/rover/commit/e62a58302e17c7ca9739349b60fc654cce7f7da3))
* use mdns for tests ([bb59b70](https://github.com/VU-ASE/rover/commit/bb59b70cbbbc7e8a31fdf4e93a85f63346c22e5f))
* using mv instead of os.Rename ([696ecc4](https://github.com/VU-ASE/rover/commit/696ecc440505fb39c3ec5cf814b55786b9fb5bf5))
* version mismatch checking ([4ba63b9](https://github.com/VU-ASE/rover/commit/4ba63b901f88055e68db2a7234ca17659588798c))
* version update script now accepts "v" prefix ([362c24c](https://github.com/VU-ASE/rover/commit/362c24c0b3264ce7d3aef7fdfabbfb6bea72d45a))
* version utility and correct comparison ([9251512](https://github.com/VU-ASE/rover/commit/92515124a450c2edd2ad8aed10a20b7b42ba4b9c))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))
* release 0.0.1 ([abd806e](https://github.com/VU-ASE/rover/commit/abd806e59353ee5ea9da82a80612ff9ece653285))
* release 0.0.5 ([aa2ed11](https://github.com/VU-ASE/rover/commit/aa2ed11e18453f8a6edc3f82cd503e38288e7060))
* release 1.0.0 ([55d09fd](https://github.com/VU-ASE/rover/commit/55d09fd29d0bc18c528ce63af45a2c10c3428f8d))

## [1.0.0](https://github.com/VU-ASE/rover/compare/v0.21.0...v1.0.0) (2025-03-31)


### ⚠ BREAKING CHANGES

* release rover v1.0.0

### Features

* release rover v1.0.0 ([292c229](https://github.com/VU-ASE/rover/commit/292c22939aeb4b8592c2dbe8ac1566c350d83b1a))

## [0.21.0](https://github.com/VU-ASE/rover/compare/v0.20.0...v0.21.0) (2025-03-31)


### Features

* unified toasts ([ded557b](https://github.com/VU-ASE/rover/commit/ded557be905b07b688a2fe9f6098a028543ec8bb))


### Bug Fixes

* auto-update roverd in test ([9b53863](https://github.com/VU-ASE/rover/commit/9b53863f55a2baf712b25f46099b1d2bbd7b7ff3))
* do not reset calibration value for normal pipeline testing ([be21d46](https://github.com/VU-ASE/rover/commit/be21d462d992f50a90e93b473417e6a0f026b3f3))
* do not show "shortcuts" tab as active on debugpage ([822b084](https://github.com/VU-ASE/rover/commit/822b084f1a15a00c34a12a162d8c618e61bc5c65))
* do NOT update roverd in the test script ([b5f1cb0](https://github.com/VU-ASE/rover/commit/b5f1cb0ea34f527117f6aee6a8d071002786c52d))
* navbar type error ([d66c700](https://github.com/VU-ASE/rover/commit/d66c70009adc0a222eba5967af53d168de2f44af))
* use mdns for tests ([bb59b70](https://github.com/VU-ASE/rover/commit/bb59b70cbbbc7e8a31fdf4e93a85f63346c22e5f))

## [0.20.0](https://github.com/VU-ASE/rover/compare/v0.19.0...v0.20.0) (2025-03-31)


### Features

* helpful toasts and warning messages when pipeline starts/stops ([27074f6](https://github.com/VU-ASE/rover/commit/27074f67ca459497ca9105b6ec2c419f70a2d49b))
* reset calibration values on trimming ([ec8eb15](https://github.com/VU-ASE/rover/commit/ec8eb15b7efc09a2b245a1424f21fb4bb51beb9d))
* show transceiver in roverctl-web ([be2cb92](https://github.com/VU-ASE/rover/commit/be2cb9274d7bed0c2d4e890b33018c8f82fbd652))


### Bug Fixes

* more lenient timeouts for roverd downloads ([a89b302](https://github.com/VU-ASE/rover/commit/a89b302fd97b48ce176d1f27becee737a67ac803))
* roverd shutsdown pipeline on termination ([83ea1dc](https://github.com/VU-ASE/rover/commit/83ea1dce347970d9a265472b62a18031177b46eb))

## [0.19.0](https://github.com/VU-ASE/rover/compare/v0.18.4...v0.19.0) (2025-03-31)


### Features

* add shortcuts item in navbar ([5aea4d3](https://github.com/VU-ASE/rover/commit/5aea4d3c7ec4458cd4aa6feb79e62398d5c1d860))
* use mDNS to resolve hostnames, allow proxy IP override ([2cb436e](https://github.com/VU-ASE/rover/commit/2cb436e2038e2ef1e5d722913f0d794ce9d60bfe))

## [0.18.4](https://github.com/VU-ASE/rover/compare/v0.18.3...v0.18.4) (2025-03-30)


### Bug Fixes

* correct link if upload fails ([d229edb](https://github.com/VU-ASE/rover/commit/d229edbeca49c5684dd881e97dc9d02bf9cbb9fc))

## [0.18.3](https://github.com/VU-ASE/rover/compare/v0.18.2...v0.18.3) (2025-03-28)


### Bug Fixes

* added service-template for cpp ([cf5fd0c](https://github.com/VU-ASE/rover/commit/cf5fd0c04b646dfef25138bbbb54220c03a2722e))

## [0.18.2](https://github.com/VU-ASE/rover/compare/v0.18.1...v0.18.2) (2025-03-27)


### Bug Fixes

* emergency stop ([2b7bdf8](https://github.com/VU-ASE/rover/commit/2b7bdf85012d4b5e736a5bf3543a7f5856bc9855))
* roverctl --force warning message ([5384c70](https://github.com/VU-ASE/rover/commit/5384c704cac253524bc7d433d4b260602e1c3632))

## [0.18.1](https://github.com/VU-ASE/rover/compare/v0.18.0...v0.18.1) (2025-03-26)


### Bug Fixes

* infinite redirect ([0f63b01](https://github.com/VU-ASE/rover/commit/0f63b019402c498c496f4cec1bbf9f483de3065a))

## [0.18.0](https://github.com/VU-ASE/rover/compare/v0.17.0...v0.18.0) (2025-03-26)


### Features

* keyboard shortcuts for emergency, stopping pipelin and debug pause ([ad50f10](https://github.com/VU-ASE/rover/commit/ad50f1014705f1631c973928ff06922c74733514))


### Bug Fixes

* incorrect number of arguments for service init script ([49981a0](https://github.com/VU-ASE/rover/commit/49981a094084063c1f584caf0f3df2a25e13aeca))
* merge service name and service "as" name in tuning overview ([3675c0e](https://github.com/VU-ASE/rover/commit/3675c0e5159f0e503c8ff65a596c5f757263a9c5))
* proxy server deadlock on repeated transceiver crashes ([577a1a2](https://github.com/VU-ASE/rover/commit/577a1a29775debc7e75526abbc9f799e9b3db0ec))
* retain debug toggle information and don't show unusable tuning ([9d78947](https://github.com/VU-ASE/rover/commit/9d789474196bc613cea17b1da0a775e8669e9d12))
* show service crash exit code on transition from started to stopped ([e43cb26](https://github.com/VU-ASE/rover/commit/e43cb260ef1df98869502afda7731fe188257beb))

## [0.17.0](https://github.com/VU-ASE/rover/compare/v0.16.0...v0.17.0) (2025-03-26)


### Features

* emergency command ([2b883cd](https://github.com/VU-ASE/rover/commit/2b883cd7ded78acf211950161ec854884722badd))
* install script with symlink instead of alias ([22ccd67](https://github.com/VU-ASE/rover/commit/22ccd6786c112cbde2979064db8ae6ef89fd5c46))
* shutdown command ([4a124b7](https://github.com/VU-ASE/rover/commit/4a124b7ddec97a31e733b064745a4374e4cd71fc))

## [0.16.0](https://github.com/VU-ASE/rover/compare/v0.15.0...v0.16.0) (2025-03-26)


### Features

* add alias "rover" to roverctl install script ([1ec8389](https://github.com/VU-ASE/rover/commit/1ec83893a0a7458aef2c8b41791c548ffda961df))
* add incompatibility warning to all endpoints that need it ([126849c](https://github.com/VU-ASE/rover/commit/126849cd89e743d089f2765f7261df24c30fd0eb))
* add installation location to service info view ([1a0ca7c](https://github.com/VU-ASE/rover/commit/1a0ca7c58c0a9b58ef3a0073babc92470ce99fe1))
* build command ([ea610eb](https://github.com/VU-ASE/rover/commit/ea610eb0171be57fbfd41ca2991bc6a1d00dedea))
* improved error printing with multiline support ([8c64c2c](https://github.com/VU-ASE/rover/commit/8c64c2ce5443774ab1571658d5f57e932cb426b1))
* improved error reporting on upload failure ([b657f17](https://github.com/VU-ASE/rover/commit/b657f17d60bdaf464a2e998ddbb1541fead96141))
* info command ([478704e](https://github.com/VU-ASE/rover/commit/478704ec6fa6d5998f16da28bf9116df0c3dac45))
* service delete endpoint ([185362f](https://github.com/VU-ASE/rover/commit/185362f0f91d3af1cac2b0abe58e4741260740f9))
* shorthand aliasses for most commands ([3c270ce](https://github.com/VU-ASE/rover/commit/3c270ce7e0fbff41f1273a2cedd927f88d24fb99))


### Bug Fixes

* display warning if author is not set in info command ([b2c5eaa](https://github.com/VU-ASE/rover/commit/b2c5eaa5e866c87066cdd2eddc1aecd6f680d493))

## [0.15.0](https://github.com/VU-ASE/rover/compare/v0.14.0...v0.15.0) (2025-03-25)


### Features

* calibration command for actuator service ([a549e94](https://github.com/VU-ASE/rover/commit/a549e9499f9b629e89b40ea6db1fc4e1e6f025c4))
* install command ([955fea9](https://github.com/VU-ASE/rover/commit/955fea970534b51d0e8e31abef62600f776475c1))

## [0.14.0](https://github.com/VU-ASE/rover/compare/v0.13.0...v0.14.0) (2025-03-25)


### Features

* pin version dependency between roverctl-roverd and warn if broken ([c8e98dc](https://github.com/VU-ASE/rover/commit/c8e98dcc4b89d079625197e91685f5b085be4bfd))


### Bug Fixes

* image push only succeeds on --verbose flag, port checking ([b0d5998](https://github.com/VU-ASE/rover/commit/b0d599895c6cc1f3365eca1aee7ceadbadb6e13e))
* show correct fix for updating roverd on mismatch ([a49144f](https://github.com/VU-ASE/rover/commit/a49144feec173a050d84a6bfee482bcb7fe7afc3))

## [0.13.0](https://github.com/VU-ASE/rover/compare/v0.12.1...v0.13.0) (2025-03-21)


### Features

* error types for the configuration endpoint ([016a94b](https://github.com/VU-ASE/rover/commit/016a94bd55a8643d300bf94d1c88c54bdb60512b))
* implemented in place update API ([44a8c73](https://github.com/VU-ASE/rover/commit/44a8c73ffd2ad93d22056335e78b3a8ff4a5da39))


### Bug Fixes

* gitignore was missing readme ([44c64ff](https://github.com/VU-ASE/rover/commit/44c64ffed3c5d317704599bfb95af5ace86af280))

## [0.12.1](https://github.com/VU-ASE/rover/compare/v0.12.0...v0.12.1) (2025-03-20)


### Bug Fixes

* transceiver output address should use `*` instead of `localhost` ([19264da](https://github.com/VU-ASE/rover/commit/19264dac0d2c18e4ff8485889e4f9dd0dd0fddab))

## [0.12.0](https://github.com/VU-ASE/rover/compare/v0.11.3...v0.12.0) (2025-03-13)


### Features

* link source code and docs for official services ([3ac9131](https://github.com/VU-ASE/rover/commit/3ac9131e4fce750e5475ec18167cf9feb83109c1))


### Bug Fixes

* correct timing of service error messages to actual moment of crash ([68f158f](https://github.com/VU-ASE/rover/commit/68f158f9404125399fc18499d855f04f6ab054f0))

## [0.11.3](https://github.com/VU-ASE/rover/compare/v0.11.2...v0.11.3) (2025-03-13)


### Bug Fixes

* roverd not propagating stream outputs to transceiver if no consumer ([0818f38](https://github.com/VU-ASE/rover/commit/0818f38c0b4353bb5078a3c8b904be8a9e5f3c38))

## [0.11.2](https://github.com/VU-ASE/rover/compare/v0.11.1...v0.11.2) (2025-03-13)


### Bug Fixes

* pipeline not loading after viewing debug screen ([67787c4](https://github.com/VU-ASE/rover/commit/67787c480eb8d8c92381433008e21b2297b5b6eb))

## [0.11.1](https://github.com/VU-ASE/rover/compare/v0.11.0...v0.11.1) (2025-03-13)


### Bug Fixes

* do not error on existing connection being reused ([6e357bf](https://github.com/VU-ASE/rover/commit/6e357bfd6dc7c008d09f0717992a3619f7f2240c))

## [0.11.0](https://github.com/VU-ASE/rover/compare/v0.10.2...v0.11.0) (2025-03-13)


### Features

* show which service crashed the pipeline ([8591295](https://github.com/VU-ASE/rover/commit/85912954080cdb6654596702b195fb1c2b064ded))


### Bug Fixes

* disable transceiver service when debug mode is not enabled ([46f41e5](https://github.com/VU-ASE/rover/commit/46f41e5f75025b3dd3515852eb42f69f88afccea))
* improved upload description and example usage ([248002c](https://github.com/VU-ASE/rover/commit/248002c3d1c96494a91577621a270ea0a3a68cea))
* load list of services from roverd in debug window ([a8099bb](https://github.com/VU-ASE/rover/commit/a8099bb929543b3dfde14b06bc61115c880f3ceb))
* more helpful messages to enable debug mode ([106e3ea](https://github.com/VU-ASE/rover/commit/106e3ea929bc7943948e77f07074171e4a31ad55))
* only show services when the pipeline is started (in debug mode) ([e22facf](https://github.com/VU-ASE/rover/commit/e22facf86deb8c7b305238dcc704b6aab7fc741b))
* service logs cut-off point makes them unreadable ([412a8ee](https://github.com/VU-ASE/rover/commit/412a8ee5eb7520daee8fbef481b0271eadab9470))
* show logs of last run when selected ([0cca2da](https://github.com/VU-ASE/rover/commit/0cca2dab6e762f6fb8916ed3cce3175f17262ec0))
* show pipeline edges accurately ([ff53e34](https://github.com/VU-ASE/rover/commit/ff53e3424e6c8644bd943c3d181c080cf69f6a74))
* sort services alphabetically. Always show vu-ase first ([fb6a58b](https://github.com/VU-ASE/rover/commit/fb6a58b9419af46bb1b06792e548bdc9431fb6da))

## [0.10.2](https://github.com/VU-ASE/rover/compare/v0.10.1...v0.10.2) (2025-03-12)


### Bug Fixes

* roverd shows all exit codes now ([2de0226](https://github.com/VU-ASE/rover/commit/2de022691173795eac6a63eec9787d554f466958))
* using mv instead of os.Rename ([696ecc4](https://github.com/VU-ASE/rover/commit/696ecc440505fb39c3ec5cf814b55786b9fb5bf5))

## [0.10.1](https://github.com/VU-ASE/rover/compare/v0.10.0...v0.10.1) (2025-03-11)


### Bug Fixes

* specifying value types in service.yaml is required ([e8757a3](https://github.com/VU-ASE/rover/commit/e8757a3edeb67edcccb0cd5d6909d15879d6f886))

## [0.10.0](https://github.com/VU-ASE/rover/compare/v0.9.0...v0.10.0) (2025-03-08)


### Features

* roverctl service init endpoint ([b8a9027](https://github.com/VU-ASE/rover/commit/b8a90274bc33f681fc5eaf97cb8adf04b828122f))

## [0.9.0](https://github.com/VU-ASE/rover/compare/v0.8.0...v0.9.0) (2025-03-08)


### Features

* author command ([6ead41c](https://github.com/VU-ASE/rover/commit/6ead41cdbfb58ab0ce484cb627ae01bae0223da1))
* logs endpoint ([f1b3077](https://github.com/VU-ASE/rover/commit/f1b3077fb46bcc396dc9a7c1d32c96b8a66537d7))
* pipeline enable/disable endpoints ([148fb7e](https://github.com/VU-ASE/rover/commit/148fb7eb16da42ae3af8ce568a767419260d9704))
* pipeline endpoint ([d0beca5](https://github.com/VU-ASE/rover/commit/d0beca50b2d01d0cedaada7833edd6a8bcba1053))
* services endpoint ([a5e11a4](https://github.com/VU-ASE/rover/commit/a5e11a46ca7d8c5a20345a356f32943e86ba50f0))
* update endpoints ([e1422b4](https://github.com/VU-ASE/rover/commit/e1422b40ed2e25e98eb094bc088154d896a28caf))


### Bug Fixes

* roverd log lines ([97adc00](https://github.com/VU-ASE/rover/commit/97adc0044f551ce2aa9ea1e7db9beeae24e43ab7))
* service yamls can now specify a bash command for run ([6022928](https://github.com/VU-ASE/rover/commit/60229285d294b95698faef9958704e3d8c81ff5e))

## [0.8.0](https://github.com/VU-ASE/rover/compare/v0.7.2...v0.8.0) (2025-03-08)


### Features

* merge passthrough into roverctl, debugging now works ([8f6c88f](https://github.com/VU-ASE/rover/commit/8f6c88f43323856971c7f6f129c47a0d23ba3a1e))


### Bug Fixes

* remove unused github action jobs ([04ef005](https://github.com/VU-ASE/rover/commit/04ef0051fa635a52e6fbc521c6ca915b58c1423e))

## [0.7.2](https://github.com/VU-ASE/rover/compare/v0.7.1...v0.7.2) (2025-03-04)


### Bug Fixes

* hard-coded version ([058573d](https://github.com/VU-ASE/rover/commit/058573d9cf67a814ef554712b4ad610eaabcb0b4))

## [0.7.1](https://github.com/VU-ASE/rover/compare/v0.7.0...v0.7.1) (2025-03-04)


### Bug Fixes

* roverd early check for root ([451f7eb](https://github.com/VU-ASE/rover/commit/451f7eb4ab031fe82dc0d7b19e2a7377122a2e5f))

## [0.7.0](https://github.com/VU-ASE/rover/compare/v0.6.0...v0.7.0) (2025-03-03)


### Features

* CLI approach to roverctl ([bb78304](https://github.com/VU-ASE/rover/commit/bb78304cf2c6fcda4b2ce95bba51fcd183ee6ce3))
* dangers page ([a296ad3](https://github.com/VU-ASE/rover/commit/a296ad3ccb2d682d7a881c057d5a60aec36a9888))
* info target ([0bc9cb0](https://github.com/VU-ASE/rover/commit/0bc9cb0ecfd0d6ffe2aac31c12717839d7ede73f))
* roverctl-web startup through docker ([e3ca9a5](https://github.com/VU-ASE/rover/commit/e3ca9a54be65233a5f7a8f8fa8b915c9d042b724))
* start roverctl-web container ([9329015](https://github.com/VU-ASE/rover/commit/9329015b9cb86a8db86bc756052327fdf170785e))
* update target ([0f3cc2a](https://github.com/VU-ASE/rover/commit/0f3cc2a340b5a78a3c04fc363941608fa9622613))

## [0.6.0](https://github.com/VU-ASE/rover/compare/v0.5.1...v0.6.0) (2025-03-03)


### Features

* add passthrough back ([e07dc70](https://github.com/VU-ASE/rover/commit/e07dc7073c6b63dd9ed1c9894a11df2127e7d23c))


### Bug Fixes

* rename passthrough to appropriate folder ([5c1d41e](https://github.com/VU-ASE/rover/commit/5c1d41e8f3a1bb3be425841ede4c9a80fe2c2861))
* roverd unclosed delimiter ([17d8e95](https://github.com/VU-ASE/rover/commit/17d8e95a9b322a205ffe46c5dd367459f59c9aef))

## [0.5.1](https://github.com/VU-ASE/rover/compare/v0.5.0...v0.5.1) (2025-03-03)


### Bug Fixes

* use lowercase for GHCR tag ([e62a583](https://github.com/VU-ASE/rover/commit/e62a58302e17c7ca9739349b60fc654cce7f7da3))

## [0.5.0](https://github.com/VU-ASE/rover/compare/v0.4.1...v0.5.0) (2025-03-03)


### Features

* added Dockerfile for roverctl-web ([ec9733b](https://github.com/VU-ASE/rover/commit/ec9733b2f507d03d0ec405e6bcc2d82519406e6b))
* added passthrough and debugging functionality ([90975ef](https://github.com/VU-ASE/rover/commit/90975ef34e1acf2ffc1f9ec740e18b600df7d561))
* app config and overview page ([1c76d8c](https://github.com/VU-ASE/rover/commit/1c76d8cd42abe86a77e5133806ac97921a4da5d5))
* enable and install debug logic from roverctl-web ([a760e30](https://github.com/VU-ASE/rover/commit/a760e30d6991e0322399aeb4d65beca78a06dcc6))
* github actions workflows to build and push docker images ([5c4b035](https://github.com/VU-ASE/rover/commit/5c4b035329c9998331228113913a28fcc72b7552))
* manage page (wip) and status overlay ([e9132ed](https://github.com/VU-ASE/rover/commit/e9132ed41535b49fbc10353c023d6935edf82640))
* service information ([c0119a5](https://github.com/VU-ASE/rover/commit/c0119a5c4715cb39d86fdecdf856ce3d48127595))
* service installation ([246a452](https://github.com/VU-ASE/rover/commit/246a452e0ae9b4d8489b1f852d1a425c7137e895))
* service selection toggles and ordening ([6ad02f4](https://github.com/VU-ASE/rover/commit/6ad02f49222b35caec422dc1e2c486276a50aa65))
* starting, stopping and modifying pipelines ([e54a176](https://github.com/VU-ASE/rover/commit/e54a1762fdc686ec0236874c1169ea9c20014bff))


### Bug Fixes

* consistent style for services ([a9f5fbe](https://github.com/VU-ASE/rover/commit/a9f5fbef311b50ce30ca3891a9c5591ec3c12d91))
* poll pipeline status ([a02bbf6](https://github.com/VU-ASE/rover/commit/a02bbf698dbd8888b4f8efcf82dfb2fdb7eca9a8))
* returning service_as optimistically for all endpoints ([a019fe1](https://github.com/VU-ASE/rover/commit/a019fe1bee5a0a5e591911b389ee599b647e0440))
* show correct panels for selected services ([5ba569e](https://github.com/VU-ASE/rover/commit/5ba569e06f205686fd5baa2dd7ff0ae765d06bb7))
* show debug toggle correctly ([3ea63f8](https://github.com/VU-ASE/rover/commit/3ea63f8a79ee4fe7ca0836407fb96d6f501fd94d))
* use dynamic env variables for configuration ([25d5c88](https://github.com/VU-ASE/rover/commit/25d5c88b217c9040aea12686a4b08b05443a8e03))

## [0.4.1](https://github.com/VU-ASE/rover/compare/v0.4.0...v0.4.1) (2025-02-12)


### Bug Fixes

* version utility and correct comparison ([9251512](https://github.com/VU-ASE/rover/commit/92515124a450c2edd2ad8aed10a20b7b42ba4b9c))

## [0.4.0](https://github.com/VU-ASE/rover/compare/v0.3.3...v0.4.0) (2025-02-11)


### Features

* service_as feature ([782ab78](https://github.com/VU-ASE/rover/commit/782ab7874c1eb6eec0e58ac73283f2f9ab940a5b))


### Bug Fixes

* hide updater screen when all services are installed ([f4dce16](https://github.com/VU-ASE/rover/commit/f4dce16e2ae91666336fa361d6c2d15db813c380))

## [0.3.3](https://github.com/VU-ASE/rover/compare/v0.3.2...v0.3.3) (2025-02-10)


### Bug Fixes

* update checking for default pipeline ([3f67220](https://github.com/VU-ASE/rover/commit/3f672204d7faaeae6ec4a3d2ec141ab52af4dc55))

## [0.3.2](https://github.com/VU-ASE/rover/compare/v0.3.1...v0.3.2) (2025-02-10)


### Bug Fixes

* version mismatch checking ([4ba63b9](https://github.com/VU-ASE/rover/commit/4ba63b901f88055e68db2a7234ca17659588798c))

## [0.3.1](https://github.com/VU-ASE/rover/compare/v0.3.0...v0.3.1) (2025-02-09)


### Bug Fixes

* roverctl replace FQN types ([f6a9c44](https://github.com/VU-ASE/rover/commit/f6a9c44dee576b4ba63f3e44ca5a07704ead245b))

## [0.3.0](https://github.com/VU-ASE/rover/compare/v0.2.0...v0.3.0) (2025-02-09)


### Features

* api adaptation ([d36262f](https://github.com/VU-ASE/rover/commit/d36262f0e0d336e08cc1a3a6e2484ad4e52c2431))
* default ASE pipeline updates and installs ([0abcdb0](https://github.com/VU-ASE/rover/commit/0abcdb0021d0c2c4158a52e479684615864518af))


### Bug Fixes

* make roverctl compatible with new API spec ([6eb675c](https://github.com/VU-ASE/rover/commit/6eb675c87be7b49d31ca3405906767df359405f3))

## [0.2.0](https://github.com/VU-ASE/rover/compare/v0.1.2...v0.2.0) (2025-02-07)


### Features

* api clean up and addition of alias error (BBC) ([1d685a6](https://github.com/VU-ASE/rover/commit/1d685a6578ddc0b9c3b48fa9ffeb0abecfefad22))
* install script that allows for installing a specific version ([f5bfcaf](https://github.com/VU-ASE/rover/commit/f5bfcaf2bada29e9e8732b61d79c93a53c3331d9))
* upgrade roverctl to go 1.22.12 ([6a19635](https://github.com/VU-ASE/rover/commit/6a1963502affb37a09328151d1c55622ffb14d0f))


### Bug Fixes

* error on roverctl install.sh if release does not exist or no binary ([1880064](https://github.com/VU-ASE/rover/commit/1880064e54514a502ae07bd21f690b8d953a7eea))
* roverctl install.sh, verify downloaded binary is valid ([381a5ed](https://github.com/VU-ASE/rover/commit/381a5ed2e809710a882a65e4affd3fb658ecb2e2))

## [0.1.2](https://github.com/VU-ASE/rover/compare/v0.1.1...v0.1.2) (2025-02-06)


### Bug Fixes

* version update script now accepts "v" prefix ([362c24c](https://github.com/VU-ASE/rover/commit/362c24c0b3264ce7d3aef7fdfabbfb6bea72d45a))

## [0.1.1](https://github.com/VU-ASE/rover/compare/v0.1.0...v0.1.1) (2025-02-06)


### Bug Fixes

* testing ci again ([573bb0f](https://github.com/VU-ASE/rover/commit/573bb0f5aafb6ca7b18208a4c71b1a14724de484))

## [0.1.0](https://github.com/VU-ASE/rover/compare/v0.0.6...v0.1.0) (2025-02-06)


### Features

* added manage script ([e1b1429](https://github.com/VU-ASE/rover/commit/e1b142989b1b0c0dfcaa55f0a496da56031288e4))


### Bug Fixes

* broken links ([7a4d2aa](https://github.com/VU-ASE/rover/commit/7a4d2aa8ee9bc87d5f06ea3c1f55a9f324cc57a7))
* ci testing ([269b7a3](https://github.com/VU-ASE/rover/commit/269b7a3b9ece651b79b22145db9eccb1cf018fa9))
* docs and install script for roverctl ([a069205](https://github.com/VU-ASE/rover/commit/a069205b726187a5ad7ad99c8f00a75f7569ce9f))
* public link for schematic in docs ([b26decf](https://github.com/VU-ASE/rover/commit/b26decf57eb08b15e0fd4c2c04e81d1f9b3a1fdd))

## [0.0.6](https://github.com/VU-ASE/rover/compare/v0.0.5...v0.0.6) (2025-02-06)


### Bug Fixes

* point install scripts to rover repository (was roverctl) ([b139324](https://github.com/VU-ASE/rover/commit/b139324d303d27c88f2f8d91d0aaf1b78a5f714a))

## [0.0.5](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.5) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))
* release 0.0.5 ([aa2ed11](https://github.com/VU-ASE/rover/commit/aa2ed11e18453f8a6edc3f82cd503e38288e7060))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.1) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* remove whitespace from upload action ([0b89220](https://github.com/VU-ASE/rover/commit/0b892200cb4e12c3321ec080a40dc8286951ac3c))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.1) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))
* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))
* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v0.0.3...v0.0.1) (2025-02-06)


### Bug Fixes

* exact match with roverd upload ([b8d38a6](https://github.com/VU-ASE/rover/commit/b8d38a6c927bc1539bece3833dff9ece119b07e6))


### Miscellaneous Chores

* release 0.0.1 ([0112285](https://github.com/VU-ASE/rover/commit/0112285b31119bb0d0262a3796a91415932d9c3e))

## [0.0.3](https://github.com/VU-ASE/rover/compare/v0.0.2...v0.0.3) (2025-02-06)


### Bug Fixes

* use correct file separator for asset upload ([699d913](https://github.com/VU-ASE/rover/commit/699d913477799a2e487821df38645650c3d5b7b5))

## [0.0.2](https://github.com/VU-ASE/rover/compare/v0.0.1...v0.0.2) (2025-02-06)


### Bug Fixes

* debug missing assets ([508f3dc](https://github.com/VU-ASE/rover/commit/508f3dcef4aab0035d5709d8f9a3922f5a1c26e6))

## [0.0.1](https://github.com/VU-ASE/rover/compare/v1.0.14...v0.0.1) (2025-02-06)


### Miscellaneous Chores

* release 0.0.1 ([abd806e](https://github.com/VU-ASE/rover/commit/abd806e59353ee5ea9da82a80612ff9ece653285))

## [1.0.14](https://github.com/VU-ASE/rover/compare/v1.0.13...v1.0.14) (2025-02-06)


### Bug Fixes

* extract from correct location ([d4106ba](https://github.com/VU-ASE/rover/commit/d4106ba17d7e8ec4d7fea3f3827aded9d4a25d0a))

## [1.0.13](https://github.com/VU-ASE/rover/compare/v1.0.12...v1.0.13) (2025-02-06)


### Bug Fixes

* extract artifact downloads first ([37762b7](https://github.com/VU-ASE/rover/commit/37762b77e75becc3a8fb960e83ddd72f47305729))

## [1.0.12](https://github.com/VU-ASE/rover/compare/v1.0.11...v1.0.12) (2025-02-06)


### Bug Fixes

* use correct version of action/artifact ([300beda](https://github.com/VU-ASE/rover/commit/300beda2031241877d4101279e5f9ec5cf8dcd07))

## [1.0.11](https://github.com/VU-ASE/rover/compare/v1.0.10...v1.0.11) (2025-02-06)


### Bug Fixes

* release assets on all success ([4448019](https://github.com/VU-ASE/rover/commit/44480199fcac3cc9c81e317ef7b34c270641ec6a))

## [1.0.10](https://github.com/VU-ASE/rover/compare/v1.0.9...v1.0.10) (2025-02-06)


### Bug Fixes

* remove unused view ([31f91cd](https://github.com/VU-ASE/rover/commit/31f91cd1bd482a2a1bb05af107ebb19b153d1de5))

## [1.0.9](https://github.com/VU-ASE/rover/compare/v1.0.8...v1.0.9) (2025-02-06)


### Bug Fixes

* permission issues for building ([8f82d02](https://github.com/VU-ASE/rover/commit/8f82d02b4b4bcd6a5930c5faa8fd504042a69f0b))

## [1.0.8](https://github.com/VU-ASE/rover/compare/v1.0.7...v1.0.8) (2025-02-06)


### Bug Fixes

* build to ./bin ([dba0f2b](https://github.com/VU-ASE/rover/commit/dba0f2b1cfd4deaa257beb7c0bbe241b02879f5d))

## [1.0.7](https://github.com/VU-ASE/rover/compare/v1.0.6...v1.0.7) (2025-02-06)


### Bug Fixes

* permissions on github actions ([0b7225a](https://github.com/VU-ASE/rover/commit/0b7225ae596dd3d1e1a29a46ad26b2f6e256d9dd))

## [1.0.6](https://github.com/VU-ASE/rover/compare/v1.0.5...v1.0.6) (2025-02-06)


### Bug Fixes

* create openapi directory if it does not exist ([1f881b8](https://github.com/VU-ASE/rover/commit/1f881b8846f38ceea597ab31be031ae7c60ecf96))

## [1.0.5](https://github.com/VU-ASE/rover/compare/v1.0.4...v1.0.5) (2025-02-06)


### Bug Fixes

* use docker buildx for builds ([aa4893b](https://github.com/VU-ASE/rover/commit/aa4893b9e97492b82198164711f7cca2e534f094))

## [1.0.4](https://github.com/VU-ASE/rover/compare/v1.0.3...v1.0.4) (2025-02-05)


### Bug Fixes

* run open-api generator as devuser ([e8e84c8](https://github.com/VU-ASE/rover/commit/e8e84c85578d578299798f469782f04e9acff0f0))

## [1.0.3](https://github.com/VU-ASE/rover/compare/v1.0.2...v1.0.3) (2025-02-05)


### Bug Fixes

* open-api-gen debugging ([013bce9](https://github.com/VU-ASE/rover/commit/013bce9f4a0838ccf91ddb53e429d27cb7132178))

## [1.0.2](https://github.com/VU-ASE/rover/compare/v1.0.1...v1.0.2) (2025-02-05)


### Bug Fixes

* do not build roverd on self-hosted ([af62651](https://github.com/VU-ASE/rover/commit/af626516335e3f1c2a5682aaa8fb54b5f72deb5c))
* do not publish openapi files ([f7f7ff0](https://github.com/VU-ASE/rover/commit/f7f7ff039253c5ebfc893fa0d53b5daad304968d))
* roverctl, build openapi files in Dockerfile ([af01eb6](https://github.com/VU-ASE/rover/commit/af01eb620b565518860cb84d5665595788b0de71))

## [1.0.1](https://github.com/VU-ASE/rover/compare/v1.0.0...v1.0.1) (2025-02-05)


### Bug Fixes

* move github.com/VU-ASE/roverctl to github.com/VU-ASE/rover/roverctl ([a35e0dc](https://github.com/VU-ASE/rover/commit/a35e0dc12895750545256432284054e560e9eb65))

## 1.0.0 (2025-02-05)


### Features

* unified roverd and roverctl under 'rover' repository ([f893365](https://github.com/VU-ASE/rover/commit/f893365bda94929d5cfa78a4a5d4407ba4d87d95))
