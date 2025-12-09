# Sova

一時ファイルをエディタで編集し、入力内容をクリップボードにコピーするシンプルなGoツールです。

## Features

- **シンプル**: 余計な機能を排除したミニマルなツール
- **クロスプラットフォーム**: macOS（pbcopy）、Linux（xclip/xsel）、Windows（clip）に対応
- **カスタマイズ可能**: エディタや履歴保存先を設定ファイルでカスタマイズ
- **履歴機能**: コピーした内容を自動保存し、後から確認可能
- **自動クリーンアップ**: 一時ファイルは自動的に削除されます
- **末尾改行除去**: コピー時に末尾の改行文字を自動削除

## Installation

```bash
go install github.com/HMasataka/sova@latest
```

## Usage

### 基本的な使い方

```bash
# エディタを開いてテキストを入力し、クリップボードにコピー
sova

# 履歴を表示
sova --history
# または
sova -H

# ヘルプを表示
sova --help
```

### 設定ファイル

設定ファイルを使ってSovaの動作をカスタマイズできます。

```bash
# 設定ファイルのサンプルをコピー
cp config.example.yaml ~/.sova/config.yaml

# お好みのエディタで編集
vim ~/.sova/config.yaml
```

#### 設定例

```yaml
# エディタの指定（デフォルト: nvim）
editor: nvim

# 履歴の保存先（デフォルト: ~/.sova/history.txt）
history_path: ~/.sova/history.txt

# 履歴の最大保存件数（0で無制限、デフォルト: 0）
max_history_entries: 100
```

#### エディタの設定例

```yaml
# Vim
editor: vim

# VS Code（保存して閉じるまで待機）
editor: code --wait

# Emacs
editor: emacs

# Nano
editor: nano
```

## Requirements

- Go 1.21以上
- テキストエディタ（nvim、vim、code、emacsなど）
- クリップボードコマンド：
  - macOS: `pbcopy`（標準搭載）
  - Linux: `xclip`または`xsel`
  - Windows: `clip`（標準搭載）

## License

MIT License - see the [LICENSE](LICENSE) file for details.
