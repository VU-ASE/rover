# Changelog

> [!NOTE]  
> This changelog is pulled from the original roverctl repo (before merge into rover). This changelog is not updated anymore, see the root changelog for the rover repository instead.

## [1.8.1](https://github.com/VU-ASE/rover/roverctl/compare/v1.8.0...v1.8.1) (2025-01-23)


### Bug Fixes

* don't hide update button if no update found (yet) ([e067d30](https://github.com/VU-ASE/rover/roverctl/commit/e067d30cfa05acebfd049829f26022994a8b7ed6))

## [1.8.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.7.0...v1.8.0) (2025-01-23)


### Features

* SSH into Rover ([3e038ff](https://github.com/VU-ASE/rover/roverctl/commit/3e038fff45803d0d3eae086fb6b126ffa165ffc0))


### Bug Fixes

* constrain cursor on overview page when filtering ([26db525](https://github.com/VU-ASE/rover/roverctl/commit/26db525c721e4c3609042f6fd662ede6a01f15ff))
* update comparison without 'v' prefix ([ded8138](https://github.com/VU-ASE/rover/roverctl/commit/ded81383e785c3a7e00d8329736fc32820c10a91))

## [1.7.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.6.0...v1.7.0) (2025-01-23)


### Features

* update checking and self-updating ([f040c75](https://github.com/VU-ASE/rover/roverctl/commit/f040c752a3c85f739e6bd3fe24ff514de5d679e3))

## [1.6.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.5.0...v1.6.0) (2025-01-22)


### Features

* improved upload capabilities ([4d1678e](https://github.com/VU-ASE/rover/roverctl/commit/4d1678e71521a474da19ddbdff5de49cd2ef0751))


### Bug Fixes

* install message reffering to Rover ([4575e71](https://github.com/VU-ASE/rover/roverctl/commit/4575e71536e6c36e62f720f9f806c2647f413d65))
* shutdown battery unplug warning ([13576f3](https://github.com/VU-ASE/rover/roverctl/commit/13576f313633a9c199e19f8046f6df929cc82ad9))

## [1.5.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.4.0...v1.5.0) (2025-01-17)


### Features

* dialog to install default pipeline if not installed yet ([7c8e044](https://github.com/VU-ASE/rover/roverctl/commit/7c8e044fb29d40fb0eaefd3407f67b7980ebc667))
* dialogs renderer ([4c1fa61](https://github.com/VU-ASE/rover/roverctl/commit/4c1fa616c284089ccdde29cfb55c6264556a7f39))
* side-by-side pipeline view ([d6c905b](https://github.com/VU-ASE/rover/roverctl/commit/d6c905bdc1a4fd07133cc06dd8fc3c5161907094))


### Bug Fixes

* replace variables crash when initializing a service ([2016e74](https://github.com/VU-ASE/rover/roverctl/commit/2016e74b950219762b6c63e6b3f33cf11c5a9c87))

## [1.4.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.3.1...v1.4.0) (2025-01-17)


### Features

* remove unused views, reorganized often used views ([426002c](https://github.com/VU-ASE/rover/roverctl/commit/426002c03cee5c881c1f75a6e43cbd47d64f24d6))

## [1.3.1](https://github.com/VU-ASE/rover/roverctl/compare/v1.3.0...v1.3.1) (2025-01-17)


### Bug Fixes

* empty pipeine ([f52b3bb](https://github.com/VU-ASE/rover/roverctl/commit/f52b3bbfca3a7c4eef1e552098283ad6a69ce1ef))
* updated to work with stricter roverd ([52a50f2](https://github.com/VU-ASE/rover/roverctl/commit/52a50f2269b4c365f9d76ebbb07e1f1fd56d9a51))

## [1.3.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.2.0...v1.3.0) (2025-01-14)


### Features

* auto-save successful connections ([454b41e](https://github.com/VU-ASE/rover/roverctl/commit/454b41e927afad7de39381a536da8aed92134ee1))
* connection management ([0467ecd](https://github.com/VU-ASE/rover/roverctl/commit/0467ecd0400d206bb16cfe06b23f847ccdc4cf7b))
* connection page key bindings ([b14dcc1](https://github.com/VU-ASE/rover/roverctl/commit/b14dcc10d44d47d381865916b70332c535cacfcb))
* improved pipeline manager page ([ca8776e](https://github.com/VU-ASE/rover/roverctl/commit/ca8776ef3d25ebd29739c53860f633adcbb33602))
* initialize with author name ([1d867b3](https://github.com/VU-ASE/rover/roverctl/commit/1d867b33575473662415725c3bb5caac6a6d960e))
* overview page stable ([3a72268](https://github.com/VU-ASE/rover/roverctl/commit/3a722688908015a7293be8ff090066f9e2ec390a))
* overwrite local author to sync into own directory ([d4f285a](https://github.com/VU-ASE/rover/roverctl/commit/d4f285aa681ac58663a08986335d792bca76d129))
* shutdown rover ([f7e9b1e](https://github.com/VU-ASE/rover/roverctl/commit/f7e9b1e55c6b7eb239ca3db422dbdb93046cb4e9))
* utility to install default ASE pipeline ([b413ca9](https://github.com/VU-ASE/rover/roverctl/commit/b413ca96dec971d7d44f8b7cc3aeca69122b1150))


### Bug Fixes

* don't quit roverctl when entering 'q' in service init form ([a497306](https://github.com/VU-ASE/rover/roverctl/commit/a4973063b4b3be7503ca5383346e71e21ff09bb7))
* no longer break downloads if a repo was already initialized ([5b890b5](https://github.com/VU-ASE/rover/roverctl/commit/5b890b53b8a2f092ca310ac7c5f055362868d773))

## [1.2.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.1.0...v1.2.0) (2025-01-07)


### Features

* test new release action ([805dd70](https://github.com/VU-ASE/rover/roverctl/commit/805dd703eeacf10088f6047698fbd94f434e30c2))


### Bug Fixes

* save configuration in ~/.config/rover directory ([6730ce0](https://github.com/VU-ASE/rover/roverctl/commit/6730ce0ed81abcfd250fa206d3983b286a282389))

## [1.1.0](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.7...v1.1.0) (2024-12-30)


### Features

* compatible with service-template-go and service-template-c ([9255111](https://github.com/VU-ASE/rover/roverctl/commit/9255111b2a3714c991dc21488bda2e46fe1d528b))


### Bug Fixes

* API timeout from 60 -&gt; 600 seconds ([9dee00d](https://github.com/VU-ASE/rover/roverctl/commit/9dee00d7e4615df4478927b8cccb7bfc67c56685))
* make install script compatible with macOS BSD grep ([6443910](https://github.com/VU-ASE/rover/roverctl/commit/64439108ab140f716b03e566283371890176a2bc))

## [1.0.7](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.6...v1.0.7) (2024-12-28)


### Bug Fixes

* set version in correct package in ldflags ([7808685](https://github.com/VU-ASE/rover/roverctl/commit/780868547625937bd3fccff6a8b88e6303b3620c))

## [1.0.6](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.5...v1.0.6) (2024-12-28)


### Bug Fixes

* upload with GITHUB_TOKEN and correct permissions ([c9470a1](https://github.com/VU-ASE/rover/roverctl/commit/c9470a1d084dbcc1e635ac38034807a7ea6252f2))

## [1.0.5](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.4...v1.0.5) (2024-12-28)


### Bug Fixes

* upload with GH_PAT instead of GITHUB_TOKEN ([e5014c6](https://github.com/VU-ASE/rover/roverctl/commit/e5014c698d8f9648f3717c8c005830e78e777f28))

## [1.0.4](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.3...v1.0.4) (2024-12-28)


### Bug Fixes

* upload releases with more flexible action ([3cb0a4d](https://github.com/VU-ASE/rover/roverctl/commit/3cb0a4da54bdf33db339f110a80e9fb68978ccca))

## [1.0.3](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.2...v1.0.3) (2024-12-28)


### Bug Fixes

* upload binary assets to github directly ([13002d0](https://github.com/VU-ASE/rover/roverctl/commit/13002d05a4871900ffbc94f19ce59309b7261c5f))

## [1.0.2](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.1...v1.0.2) (2024-12-28)


### Bug Fixes

* build dir for binary output ([78a345c](https://github.com/VU-ASE/rover/roverctl/commit/78a345c7218248c1dcff61fa8c0fef52e5cd88f9))

## [1.0.1](https://github.com/VU-ASE/rover/roverctl/compare/v1.0.0...v1.0.1) (2024-12-28)


### Bug Fixes

* release binary with correct go version ([f0e3b7e](https://github.com/VU-ASE/rover/roverctl/commit/f0e3b7edb32d8608e72c399ec5c253598c2645e6))

## 1.0.0 (2024-12-28)


### Features

* add compilation workflow, proper name ([3427ee4](https://github.com/VU-ASE/rover/roverctl/commit/3427ee4dbc8f857d62225529572d7cc37fd43a1b))
* adding and testing a persistent connection ([3caafa9](https://github.com/VU-ASE/rover/roverctl/commit/3caafa96c78bb94cc4126789f82c2f12ace4e43b))
* basic connection checks ([a51875b](https://github.com/VU-ASE/rover/roverctl/commit/a51875b8747d48f37eb2f41423e3eeb548cabaf5))
* basic docs ([2e70ef8](https://github.com/VU-ASE/rover/roverctl/commit/2e70ef8cc6c304cfecf6fef0027a95b9d424943f))
* clear page separation ([5ac19fe](https://github.com/VU-ASE/rover/roverctl/commit/5ac19fecd796eacd60d75e9c75dd2c3a2c809a30))
* configure pipelines by adding/removing with validation ([db222dd](https://github.com/VU-ASE/rover/roverctl/commit/db222dd52e6e8408958584a49420073168f87242))
* connection initialization using default tui.Action interface ([b3d4854](https://github.com/VU-ASE/rover/roverctl/commit/b3d48545fa4250287ed0ec6e10b2da0cb927e4b5))
* connection setup ([36c0055](https://github.com/VU-ASE/rover/roverctl/commit/36c0055dfca45b1c248c6d1c9f5d72f6c719a94e))
* cross-platform release ([3e7d5fa](https://github.com/VU-ASE/rover/roverctl/commit/3e7d5fade853b8abb516126a3e68d04b37a4fe1e))
* enable custom hostname/IP for roverd connections ([9155a23](https://github.com/VU-ASE/rover/roverctl/commit/9155a23c39d75725fe3314d3b87d76c55927be1e))
* file watching for service synchronization ([a681412](https://github.com/VU-ASE/rover/roverctl/commit/a6814129afa5223d43a4d284b6c618ae1ee4671e))
* integrate new api spec, updated timeout and info view ([11d9140](https://github.com/VU-ASE/rover/roverctl/commit/11d91404ac742db3b036fe68dd41006a0f5dd239))
* integrate pipeline start/stop and logs ([e41b648](https://github.com/VU-ASE/rover/roverctl/commit/e41b648d458476ba56a115f8017346a3845927fc))
* list remote services ([a875e39](https://github.com/VU-ASE/rover/roverctl/commit/a875e3999a690ce5c829cc79d88d1b4c99f4c6b8))
* manual retry for failed service syncs ([10d1052](https://github.com/VU-ASE/rover/roverctl/commit/10d1052dd25b0aa5b6f28ebe56680e959f3c9ab1))
* navigation and quitting pages ([fd67a64](https://github.com/VU-ASE/rover/roverctl/commit/fd67a64aa3927afe824528bb92452764f130b3ef))
* optimistic log reloads ([674ad61](https://github.com/VU-ASE/rover/roverctl/commit/674ad61a7fcedd29bddcc2ea13f74d3b17e35fab))
* pipeline configurator ux ([8473b0a](https://github.com/VU-ASE/rover/roverctl/commit/8473b0afa13c224a3d676889be4b20eb94bc9739))
* pipeline fetch ([162322d](https://github.com/VU-ASE/rover/roverctl/commit/162322d9ca51fb7ceee64cf35105a87580ec8bb0))
* pipeline initial view ([49021ec](https://github.com/VU-ASE/rover/roverctl/commit/49021ecbef253e3b9a0a3677a8d00c269e879a12))
* pipeline overview, logs, optimistic updates, service details ([10a62f5](https://github.com/VU-ASE/rover/roverctl/commit/10a62f5af5002e10f5539c52b3a1ad58eb42eff3))
* pipeline saving and loading ([6b15717](https://github.com/VU-ASE/rover/roverctl/commit/6b157176289b46811e97735e7d9c5abe9b6c4242))
* pipeline visualization ([e4dcbe2](https://github.com/VU-ASE/rover/roverctl/commit/e4dcbe2afe7070786add445098c96f003b3bddb3))
* quality of life ([a962d5e](https://github.com/VU-ASE/rover/roverctl/commit/a962d5e0c938d915b58fa8e27fb8983bd7b1616f))
* roverd and roverctl information ([ed863cc](https://github.com/VU-ASE/rover/roverctl/commit/ed863cc33762c9317b0734c7969ca8f2a1f44646))
* save pipelines to remote ([cc32a8e](https://github.com/VU-ASE/rover/roverctl/commit/cc32a8e4bf77f599c961ee1c84f62e7ddb41a0bb))
* selected service updating ([127e09a](https://github.com/VU-ASE/rover/roverctl/commit/127e09ac7e31c9ce7043183c08e3db46840baf86))
* service list integration with roverd ([eb79b5e](https://github.com/VU-ASE/rover/roverctl/commit/eb79b5eaf1eb37cc62d814eda18c98db68a1c9e9))
* service upload through zip ([d57db41](https://github.com/VU-ASE/rover/roverctl/commit/d57db41a6843e68a5afb0988a1f5f0a0b1f6c364))
* set up open-api-generator ([6b92477](https://github.com/VU-ASE/rover/roverctl/commit/6b924775473a36234b49e94df5c69978212013df))
* shared gh actions ([d19543b](https://github.com/VU-ASE/rover/roverctl/commit/d19543b09f3d5bd1cb448ef1810397463e8dfb50))
* shortcut commands for quick access ([c42e4eb](https://github.com/VU-ASE/rover/roverctl/commit/c42e4eb035daa7586dee2214db3591fcb59249e1))
* show correct service details (including build date) ([5d3e73e](https://github.com/VU-ASE/rover/roverctl/commit/5d3e73e4189b9c34b993ccd517bb02ecc7a760b6))
* source update functionality ([d6e4fdc](https://github.com/VU-ASE/rover/roverctl/commit/d6e4fdcbd1de7b5541a934e95d937773701e2e0b))
* stable routing ([2fc2415](https://github.com/VU-ASE/rover/roverctl/commit/2fc24156bcfe53330b1c17bbf0dee6a8c6024798))
* upload services to remote ([b1912c7](https://github.com/VU-ASE/rover/roverctl/commit/b1912c77d37fcd21b0ec89e3bcd3df7493cd1d9e))
* upload with generic actions ([5f9f710](https://github.com/VU-ASE/rover/roverctl/commit/5f9f710d64e399cde7b1b84b85d15d45f886543d))
* utilities, services, connections ([b2c6db2](https://github.com/VU-ASE/rover/roverctl/commit/b2c6db2efd04a0a63819adcba1fb2901a8654c07))


### Bug Fixes

* table resizing issues ([065c52c](https://github.com/VU-ASE/rover/roverctl/commit/065c52c29147a3086032e9ebbab094e68a68bde2))

## 1.0.0 (2024-04-07)


### Features

* initial setup ([97437db](https://github.com/VU-ASE/template-GoModule/commit/97437db11b2010c16c4b13983e8740eec58431e5))
* Integrate template to servicerunner ([bbb7acd](https://github.com/VU-ASE/template-GoModule/commit/bbb7acd6a1ea5625af0902a56545a7a2e20085e6))


### Bug Fixes

* Makefile support for runargs and buildargs ([8e4c520](https://github.com/VU-ASE/template-GoModule/commit/8e4c520c12017a3c36e9cd62475a5657667b12dc))
* updated CI to zip service.yaml[D[D[D[D[D[D[D[D[D[D[D[D[D[D[D[D[D ([ff258f0](https://github.com/VU-ASE/template-GoModule/commit/ff258f0bdc202b0cb9ad5f785c377a50a0fec269))
