package common_api

import "errors"

type (
	errorMap map[int]error
)

var (
	ErrUnknownError              = errors.New("common_api: Неизвестная ошибка")
	ErrUnknownApiType            = errors.New("common_api: Неизвестный тип API")
	ErrApiDisabledInSettings     = errors.New("common_api: API отключено в настройках модуля TM API в ТМ2")
	ErrSecretKeyDoesNotMatch     = errors.New("common_api: Не совпадает секретный ключ")
	ErrUnsupportedApiVersion     = errors.New("common_api: Неподдерживаемая версия API")
	ErrUnknownRequestName        = errors.New("common_api: Неизвестное название запроса")
	ErrInvalidRequestType        = errors.New("common_api: Неверный тип запроса")
	ErrMissingParameter          = errors.New("common_api: Не хватает входного параметра")
	ErrIncorrectParameter        = errors.New("common_api: Incorrect parameter")
	ErrInternalRequestProcessing = errors.New("common_api: Внутренняя ошибка обработки запроса")

	ErrUdsNotFound                = errors.New("common_api: Служба ЕДС не найдена")
	ErrCarNotFound                = errors.New("common_api: Автомобиль не найден")
	ErrCrewNotFound               = errors.New("common_api: Экипаж не найден")
	ErrPhoneNotFound              = errors.New("common_api: Телефон не найден")
	ErrOrderNotFound              = errors.New("common_api: Заказ не найден")
	ErrTariffNotFound             = errors.New("common_api: Тариф не найден")
	ErrAccountTypeNotFound        = errors.New("common_api: Не найден тип счета")
	ErrDriverNotFound             = errors.New("common_api: Водитель не найден")
	ErrZoneNotFound               = errors.New("common_api: Район не найден")
	ErrStopNotFound               = errors.New("common_api: Стоянка не найдена")
	ErrZoneStopNotFound           = errors.New("common_api: Район остановки не найден")
	ErrZoneSourceNotFound         = errors.New("common_api: Район подачи не найден")
	ErrZoneDestinationNotFound    = errors.New("common_api: Район назначения не найден")
	ErrCoordsNotFound             = errors.New("common_api: Координаты не найдены")
	ErrClientNotFound             = errors.New("common_api: Не найден клиент")
	ErrDiscountNotFound           = errors.New("common_api: Скидка не найдена")
	ErrRouteNotRecognized         = errors.New("common_api: Маршрут не распознан")
	ErrClientGroupNotFound        = errors.New("common_api: Группа клиента не найдена")
	ErrDiscountCardNotFound       = errors.New("common_api: Дисконтная карта не найдена")
	ErrCustomerClientNotFound     = errors.New("common_api: Сотрудник клиента не найден")
	ErrDiscountCardIsNotValid     = errors.New("common_api: Дисконтная карта не действительна")
	ErrNoMatchingAddressesFound   = errors.New("common_api: Подходящие адреса не найдены")
	ErrSearchLocationNotSpecified = errors.New("common_api: Не указано место для поиска адресов")
	ErrParentClientNotFound       = errors.New("common_api: Клиент указанный в качестве родителя не найден")
	ErrUsersNotFound              = errors.New("common_api: Пользователи для отправки сообщения не найдены")
	ErrOrderStateNotFound         = errors.New("common_api: Состояние заказа не найдено")
	ErrOrderParameterNotFound     = errors.New("common_api: Параметр заказа не найден")

	ErrClientBlocked         = errors.New("common_api: Клиент заблокирован")
	ErrCustomerClientBlocked = errors.New("common_api: Сотрудник клиента заблокирован")

	ErrCashPaymentNotAllowed     = errors.New("common_api: Для клиента запрещена оплата заказа наличными. Клиент должен максимально использовать в заказе безналичную оплату (оплату с основного счета)")
	ErrNegativeBalanceCashless   = errors.New("common_api: Отрицательный баланс на безналичном счете клиента в ТМ")
	ErrInsufficientFundsCashless = errors.New("common_api: Недостаточно средств на безналичном счете клиента в ТМ")

	ErrClientwhoCanUseTheirOwnNotFound = errors.New("common_api: Не найден клиент, который может использовать собственный счет для оплаты заказов")

	ErrSpecialOrderCheck = errors.New("common_api: Ошибка специальной проверки заказа")

	ErrCalculationByTariff = errors.New("common_api: Ошибка при расчете по тарифу")

	ErrStateCannotBeChanged = errors.New("common_api: Изменение состояния не соответствует необходимым условиям")

	ErrTerminalAccountIncorrect = errors.New("common_api: Некорректный терминальный аккаунт")

	ErrCarConflictByCode               = errors.New("common_api: Автомобиль имеет такой же позывной")
	ErrCrewConflictByDriverAndCar      = errors.New("common_api: Экипаж с таким водителем и автомобилем уже существует")
	ErrDriverConflictByTerminalAccount = errors.New("common_api: Терминальный аккаунт не уникален")
	ErrDriverConflictByPrimaryPhone    = errors.New("common_api: Основной телефон может быть только один")
	ErrDriverRequiredPrimaryPhone      = errors.New("common_api: Водитель должен иметь основной телефон")
	ErrClientConflictByPhone           = errors.New("common_api: Клиент имеет такой же номер телефона")

	ErrClientExistsWithPhone    = errors.New("common_api: Клиент с таким номером телефона уже существует")
	ErrClientExistsWithLogin    = errors.New("common_api: Клиент с таким логином уже существует")
	ErrOrderExistsWithParametrs = errors.New("common_api: Заказ с такими параметрами уже создан")

	ErrForbiddenEditCrewOnLine     = errors.New("common_api: Запрещено редактирование у экипажа на линии")
	ErrUdsCarAndDriverDoesNotMatch = errors.New("common_api: Служба ЕДС автомобиля и водителя не совпадает")

	ErrAttributeNotFound           = errors.New("common_api: Атрибут не найден")
	ErrAttributeCannotBeBoundOrder = errors.New("common_api: Атрибут не может быть привязан к заказу")

	ErrParameterNotFoundOrCannotBeBoundCrew   = errors.New("common_api: Параметр не найден или не может быть привязан к экипажу")
	ErrParameterNotFoundOrCannotBeBoundCar    = errors.New("common_api: Параметр не найден или не может быть привязан к автомобилю")
	ErrParameterNotFoundOrCannotBeBoundDriver = errors.New("common_api: Параметр не найден или не может быть привязан к водителю")
)
