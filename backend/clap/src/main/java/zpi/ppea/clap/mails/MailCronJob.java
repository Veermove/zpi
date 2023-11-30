package zpi.ppea.clap.mails;

import jakarta.mail.MessagingException;
import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.io.IOException;

@Slf4j
@Component
@AllArgsConstructor
public class MailCronJob {

    private final EmailServiceImpl emailService;

    @Scheduled(cron = "${cron.job.mail}")
    public void checkEmails() {
        log.info("Sending emails");
        try {
            emailService.sendHtmlMessageFromFile();
        } catch (MessagingException | IOException e) {
            throw new RuntimeException(e);
        }
    }

}
