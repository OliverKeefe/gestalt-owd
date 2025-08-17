package net.gestalt.fileservice.adapter.persistence.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntityBase;
import jakarta.persistence.*;

import java.time.Instant;
import java.util.UUID;

@Entity
@Table(name = "files")
public class FileEntity extends PanacheEntityBase {

    @Id @GeneratedValue private UUID id;

    @GeneratedValue private UUID orgId;

    @GeneratedValue private UUID ownerId;

    @Column(name = "storage_type")
    private String storageType;

    @Column(name = "storage_key")
    private String storageKey;

    @Column(name = "file_size")
    private Long fileSize;

    @Column(name = "checksum")
    private String checksum;

    @Column(name = "created_at")
    private Instant createdAt;

    @Column(name = "last_modified_at")
    private Instant lastModifiedAt;
}
