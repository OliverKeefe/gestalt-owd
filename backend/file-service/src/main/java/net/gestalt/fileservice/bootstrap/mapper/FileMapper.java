package net.gestalt.fileservice.bootstrap.mapper;

import jakarta.enterprise.context.ApplicationScoped;
import net.gestalt.fileservice.adapter.api.dto.request.FileRequest;
import net.gestalt.fileservice.adapter.persistence.entity.FileEntity;
import net.gestalt.fileservice.core.model.FileModel;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.NullValueCheckStrategy;

@Mapper(componentModel = "cdi")
public interface FileMapper {

    FileModel toModel(FileRequest fileRequest);

    FileModel toModel(FileEntity fileEntity);
}
