mockgen -source=./pkg/repository/v1/communicationmethod/communicationmethod.go -destination=./pkg/repository/v1/communicationmethod/mock_communicationmethod/mock_communicationmethod.go
mockgen -source=./pkg/repository/v1/communicationmethodfield/communicationmethodfield.go -destination=./pkg/repository/v1/communicationmethodfield/mock_communicationmethodfield/mock_communicationmethodfield.go
mockgen -source=./pkg/repository/v1/communicationmethodlabel/communicationmethodlabel.go -destination=./pkg/repository/v1/communicationmethodlabel/mock_communicationmethodlabel/mock_communicationmethodlabel.go
mockgen -source=./pkg/repository/v1/contact/contact.go -destination=./pkg/repository/v1/contact/mock_contact/mock_contact.go
mockgen -source=./pkg/repository/v1/contactcommunicationmethod/contactcommunicationmethod.go -destination=./pkg/repository/v1/contactcommunicationmethod/mock_contactcommunicationmethod/mock_contactcommunicationmethod.go
mockgen -source=./pkg/repository/v1/contactsystem/contactsystem.go -destination=./pkg/repository/v1/contactsystem/mock_contactsystem/mock_contactsystem.go

mockgen -source=./pkg/service/v1/communicationmethodfield/communicationmethodfield.go -destination=./pkg/service/v1/communicationmethodfield/mock_communicationmethodfield/mock_communicationmethodfield.go
mockgen -source=./pkg/service/v1/communicationmethodlabel/communicationmethodlabel.go -destination=./pkg/service/v1/communicationmethodlabel/mock_communicationmethodlabel/mock_communicationmethodlabel.go
mockgen -source=./pkg/service/v1/contactsystem/contactsystem.go -destination=./pkg/service/v1/contactsystem/mock_contactsystem/mock_contactsystem.go