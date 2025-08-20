# Sova

一時ファイルをnvimで編集し、入力内容をクリップボードにコピーするシンプルなGoツールです。

## How to Use

### Install

```bash
go install github.com/HMasataka/sova
```

### Run

```bash
sova
```

## Features

- **シンプル**: 余計な機能を排除したミニマルなツール
- **クロスプラットフォーム**: macOS（pbcopy）、Linux（xclip/xsel）、Windows（clip）に対応
- **自動クリーンアップ**: 一時ファイルは自動的に削除されます
- **末尾改行除去**: コピー時に末尾の改行文字を自動削除

## Requirements

- Go 1.21以上
- nvim（テキストエディタ）
- クリップボードコマンド：
  - macOS: `pbcopy`（標準搭載）
  - Linux: `xclip`または`xsel`
  - Windows: `clip`（標準搭載）

## License

MIT License - see the [LICENSE](LICENSE) file for details.
