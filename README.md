# TheMiningS2
![Static Badge](https://img.shields.io/badge/-Skript-000000?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/-SQLite3-003B57?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/-SkBee-FFC700?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/-GraalVM-FF7800?style=for-the-badge)
![Static Badge](https://img.shields.io/badge/-Docker-2496ED?style=for-the-badge)

[![テスト](https://github.com/TeamNekozouneko/TheMiningS2/actions/workflows/test.yml/badge.svg)](https://github.com/TeamNekozouneko/TheMiningS2/actions/workflows/test.yml)

> [!NOTE]
> The Mining Season-2は、現在開発中です。  
> 本リポジトリはオープンソースであり、誰でも実行することができますが、バグや不具合がある可能性があります。

| 技術仕様 |  |
|-|-|
| バージョン | 1.21.3 |
| 開発言語 | Skript (2.10.1) |
| データベース | SQLite3 |
| アドオン | SkBee (3.8.2) |
| 実行環境 | GraalVM (Java 21) |
| テスト環境 | Github Actions（コンパイル・ランタイムテスト） |
| | Docker（ランタイム実行環境） |

## ランタイム実行環境
ディレクトリ`/`をビルドコンテキストルートとします。
### 実行（ビルド+コンテナ起動）
```bash
$ sh .runtime_env/run.sh
```
### ビルド
> [!CAUTION]
> Dockerfile内に組み込みのGraalVMの関係から`aarch64`でのみ動作します。
```bash
$ sh .runtime_env/build.sh
```
### コンテナ起動
```bash
$ sh .runtime_env/run.sh
```