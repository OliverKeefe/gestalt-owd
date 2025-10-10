package net.gestalt.app.service.token;

import io.quarkus.test.junit.QuarkusTest;
import net.gestalt.data.entity.User;
import org.junit.jupiter.api.Test;

@QuarkusTest
public class GenerateTokenTest {

    private User createMockUser() {
        User user = new User();
        user.setId(1L);
        user.setEmail("test-user@testmail.com");
        user.setSecondaryEmail("testing2@testerosa.io");

        return user;
    }

    @Test
    void shouldGenerateJwt() {
        User user = createMockUser();
    }
}
