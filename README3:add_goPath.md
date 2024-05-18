# デフォルトシェルの確認

```bash
echo $SHELL
```
結果がzshの場合は設定ファイルは.zshrc
bashの場合は.bash_profile

# 設定ファイルにパスを追加
```bash
open ~/.zshrc
```

でファイルを開き、以下を追加します。

export PATH="go env GOPATHで確認したパス/bin:$PATH"

STEP4: Pathを反映させる

```bash
source ~/.zshrc
```