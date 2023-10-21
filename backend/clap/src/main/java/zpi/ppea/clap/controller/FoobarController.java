package zpi.ppea.clap.controller;

import lombok.AllArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import zpi.ppea.clap.service.FoobarService;

@RestController
@RequestMapping("/hello")
@AllArgsConstructor
public class FoobarController {

    private final FoobarService foobarService;

    @GetMapping("/{name}")
    public ResponseEntity<String> helloGrpc(@PathVariable String name) {
        return ResponseEntity.ok(foobarService.receiveGreeting(name));
    }
}
