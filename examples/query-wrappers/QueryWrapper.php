<?php

namespace App\Database;

use PDOException;

enum LogLevelType: string
{
    case INFO = 'INFO';
    case WARN = 'WARN';
    case ERROR = 'ERROR';
    case FATAL = 'FATAL';
}

class CentralLogger
{
    public static function log(LogLevelType $level, string $message, \Throwable $error = null): void
    {
        // Central logging implementation
        error_log(sprintf("[%s] %s: %s", $level->value, $message, $error ? $error->getMessage() : ''));
    }
}

class ErrorMessages
{
    public const QUERY_FAILED = 'Database query failed';
}

class QueryResult
{
    private bool $isFail;
    private mixed $data;
    private ?string $errorMessage;

    public function __construct(bool $isFail, mixed $data, ?string $errorMessage = null)
    {
        $this->isFail = $isFail;
        $this->data = $data;
        $this->errorMessage = $errorMessage;
    }

    public function isFailed(): bool
    {
        return $this->isFail;
    }

    public function getData(): mixed
    {
        return $this->data;
    }

    public function getErrorMessage(): ?string
    {
        return $this->errorMessage;
    }
}

class QueryWrapper
{
    /**
     * Executes a PDO query, catching PDOException and returning a safe QueryResult.
     */
    public static function execute(callable $queryFn): QueryResult
    {
        try {
            $data = $queryFn();
            return new QueryResult(false, $data);
        } catch (PDOException $error) {
            CentralLogger::log(LogLevelType::ERROR, ErrorMessages::QUERY_FAILED, $error);
            return new QueryResult(true, null, ErrorMessages::QUERY_FAILED);
        }
    }
}
