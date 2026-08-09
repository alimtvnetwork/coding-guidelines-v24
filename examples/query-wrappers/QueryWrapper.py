from enum import Enum
from typing import Any, Callable, Optional, TypeVar, Generic

T = TypeVar('T')

class LogLevelType(Enum):
    INFO = 'INFO'
    WARN = 'WARN'
    ERROR = 'ERROR'
    FATAL = 'FATAL'

class CentralLogger:
    @staticmethod
    def log(level: LogLevelType, message: str, error: Optional[Exception] = None) -> None:
        print(f"[{level.value}] {message}: {error}")

class ErrorMessages(Enum):
    QUERY_FAILED = 'Database query failed'

class QueryResult(Generic[T]):
    def __init__(self, is_fail: bool, data: Optional[T] = None, error_message: Optional[str] = None):
        self.is_fail = is_fail
        self.data = data
        self.error_message = error_message

class QueryWrapper:
    @staticmethod
    def execute(query_fn: Callable[[], T]) -> QueryResult[T]:
        try:
            data = query_fn()
            return QueryResult(is_fail=False, data=data)
        except Exception as error:
            CentralLogger.log(LogLevelType.ERROR, ErrorMessages.QUERY_FAILED.value, error)
            return QueryResult(is_fail=True, data=None, error_message=ErrorMessages.QUERY_FAILED.value)
