package zpi.ppea.clap.exceptions;

import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import zpi.ppea.clap.config.ValueConfig;

import java.util.*;

@ControllerAdvice
@AllArgsConstructor
public class GlobalExceptionHandler {

    private final ValueConfig valueConfig;

    @ExceptionHandler(UserNotAuthorizedException.class)
    public ResponseEntity<String> handleUserNotAuthorizedException(UserNotAuthorizedException ex) {
        return ResponseEntity.status(HttpStatus.UNAUTHORIZED)
                .header(valueConfig.getRefreshTokenHeaderName(), ex.getRefresh()).body(ex.getMessage());
    }

    @ExceptionHandler(GrpcDatastoreException.class)
    public ResponseEntity<String> handleNoAccessToResourceException(GrpcDatastoreException ex) {
        return ResponseEntity.status(HttpStatus.FORBIDDEN)
                .header(valueConfig.getRefreshTokenHeaderName(), ex.getRefresh()).body(ex.getMessage());
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Map<String, List<String>>> handleValidationErrors(MethodArgumentNotValidException ex) {
        List<String> errors = new ArrayList<>();
        for (var messages : Objects.requireNonNull(ex.getDetailMessageArguments()))
            errors.addAll((Collection<? extends String>) messages);
        return ResponseEntity.status(HttpStatus.BAD_REQUEST)
                .header(valueConfig.getRefreshTokenHeaderName(), "false")
                .body(getErrorsMap(errors));
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<String> handleUnexpectedException(Exception ex) {
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
                .header(valueConfig.getRefreshTokenHeaderName(), "false").body(ex.getMessage());
    }

    private Map<String, List<String>> getErrorsMap(List<String> errors) {
        Map<String, List<String>> errorResponse = new HashMap<>();
        errorResponse.put("errors", errors);
        return errorResponse;
    }

}
