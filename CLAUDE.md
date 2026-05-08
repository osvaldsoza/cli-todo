
## GitHub

Repositório: https://github.com/osvaldsoza/cli-todo

Sincronização automática: um hook `PostToolUse` em `.claude/settings.local.json` faz commit e push para `origin/main` após cada edição de arquivo (ferramentas `Write` e `Edit`). Mensagem de commit gerada automaticamente com timestamp.

Para desativar temporariamente, remova ou comente o bloco `hooks` em `.claude/settings.local.json`.


