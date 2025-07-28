package net.gestalt.core.model;

import java.util.Set;

public class User {
    private Long id;
    private String email;
    private String secondaryEmail;
    private String hashedPassword;
    private Set<String> role;
    private Set<String> organization;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getSecondaryEmail() {
        return secondaryEmail;
    }

    public void setSecondaryEmail(String secondaryEmail) {
        this.secondaryEmail = secondaryEmail;
    }

    public void setHashedPassword(String hashedPassword) {
        this.hashedPassword = hashedPassword;
    }

    public String getHashedPassword() {
        return hashedPassword;
    }

    public Set<String> getOrganization() {
        return organization;
    }

    public void setOrganization(Set<String> organization) {
        this.organization = organization;
    }

    public Set<String> getRole() {
        return role;
    }

    public void setRole(Set<String> role) {
        this.role = role;
    }
}
