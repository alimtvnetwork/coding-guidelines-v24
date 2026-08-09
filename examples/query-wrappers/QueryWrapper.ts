enum LogLevelType {
  INFO = 'INFO',
  WARN = 'WARN',
  ERROR = 'ERROR',
  FATAL = 'FATAL'
}

class CentralLogger {
  static log(level: LogLevelType, message: string, error?: unknown): void {
    console.error(`[${level}] ${message}`, error);
  }
}

interface QueryResult<T> {
  data: T | null;
  isFail: boolean;
  errorMessage: string | null;
}

const ERROR_MESSAGES = {
  QUERY_FAILED: 'Database query failed'
} as const;

class QueryWrapper {
  static async execute<T>(queryFn: () => Promise<T>): Promise<QueryResult<T>> {
    try {
      const data = await queryFn();
      return {
        data,
        isFail: false,
        errorMessage: null
      };
    } catch (error) {
      CentralLogger.log(LogLevelType.ERROR, ERROR_MESSAGES.QUERY_FAILED, error);
      return {
        data: null,
        isFail: true,
        errorMessage: ERROR_MESSAGES.QUERY_FAILED
      };
    }
  }
}

export { QueryWrapper, LogLevelType, QueryResult, CentralLogger, ERROR_MESSAGES };
