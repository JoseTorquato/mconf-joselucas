import os


print('No seu navegador digite => localhost:3000')
print('Servidor rodando...')
print()
# linha de comando
os.chdir('/home/jose/vscode/teste_mconf/api_go/api')
# container docker ?
# os.chdir('/go/src')
os.system('go run server.go')