import os


print('No seu navegador digite => localhost:3000')
print('Servidor rodando...')
print()
# linha de comando localmente
os.chdir('/home/jose/vscode/teste_mconf/api_go/api')

os.system('go run server.go')