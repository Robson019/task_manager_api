package messages

const (
	AccountID       = "ID da conta do usuário"
	AccountEmail    = "E-mail da conta do usuário"
	AccountPassword = "Senha da conta do usuário"

	InvalidAccountErrorMessage         = "A conta de usuário informada é inválida!"
	InvalidAccountIDErrorMessage       = "O ID da conta do usuário informado é inválido!"
	InvalidAccountEmailErrorMessage    = "O e-mail da conta do usuário informado é inválido!"
	InvalidAccountPasswordErrorMessage = "A senha da conta do usuário informada é inválida! Ela deve ter pelo menos 8 caracteres, 1 letra maiúscula, 1 letra minúscula, 1 número e 1 caractere especial."
	InvalidCredentialsErrorMessage     = "As credenciais informadas estão incorretas!"

	NotFoundAccountErrorMessage = "A conta de usuário não foi encontrada na base de dados!"
)
