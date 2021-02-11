import os


path = './api'
print('---------------------------------')
print('Digite 1 para rodar a api em go')
print('Qualquer outra tecla para sair do programa')
print('---------------------------------')
print()
pergunta = input('O que você gostaria de fazer? ')

if pergunta == '1':
    print('No seu navegador digite => localhost:3000')
    print('Servidor rodando...')
    print()
    os.chdir('/home/jose/vscode/teste_mconf/api_go/api')
    os.system('go run server.go')