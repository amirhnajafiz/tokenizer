# Tokenizer

![GitHub top language](https://img.shields.io/github/languages/top/amirhnajafiz/tokenizer)
![GitHub last commit (by committer)](https://img.shields.io/github/last-commit/amirhnajafiz/tokenizer)
![GitHub tag checks state](https://img.shields.io/github/checks-status/amirhnajafiz/tokenizer/v0.1.0)
![GitHub release (with filter)](https://img.shields.io/github/v/release/amirhnajafiz/tokenizer)

By using ```Tokenizer``` you can manage your system tokens and credentials. You can
set your tokens, encrypt them, list them, update and delete them. The best solution
to protect your sensitive information.

## setup

Clone into repository and build ```tokenizer``` by using the following commands:

```shell
git clone https://github.com/amirhnajafiz/tokenizer.git
cd tokenizer
go build -o tokenizer
```

In order to use ```tokenizer``` in every place on your system, make sure to the followings to
either ~/.zshrc, ~/.bash_profile, or ~/.bashrc.

```shell
export PATH="<path-to-cloned-repository>:$PATH"
export TK_PATH="<path-to-cloned-repository>"
```

## encrypt

Your tokens are being stored in ```conf.txt``` file which is encrypted with
a private key that you can set by setting a value to ```TK_PRIVATE``` environment
variable. Make sure to set a 16 byte value.

```text
7nKyNrGb0Djjyv9UqT4pGaJorDA64QdsfQ==//&&//VcvhC6O4Rm3j3rmMQFMFONWnDHeXus542G4cUydlpcn98cMxyTvBI1KdcJsM
2AX/7Q/EGjMYO/wE95BESCi2o/f381KmGp63//&&//VDUD18Y4P+aBg88kfV4/pXStObtm3trslVRSprAmpHPEMZGFhwD000KISds=
```

## examples

### set

```shell
tokenizer set PWA_TOKEN "h30301j99nn968nskd[[j043jf3fj"
```

### get all

```shell
tokenizer all 

PWA_TOKEN
```

### get

```shell
tokenizer get PWA_TOKEN

h30301j99nn968nskd[[j043jf3fj
```

### delete

```shell
tokenizer del PWA_TOKEN
```
