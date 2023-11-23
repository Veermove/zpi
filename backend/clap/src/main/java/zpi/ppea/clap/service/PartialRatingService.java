package zpi.ppea.clap.service;

import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import zpi.ppea.clap.dtos.PartialRatingDto;
import zpi.ppea.clap.dtos.UpdatePartialRatingDto;
import zpi.ppea.clap.mappers.DtoMapper;
import zpi.ppea.clap.repository.PartialRatingRepository;
import zpi.ppea.clap.security.FirebaseAgent;

@Slf4j
@Service
@AllArgsConstructor
public class PartialRatingService {

    private final PartialRatingRepository partialRatingRepository;

    public PartialRatingDto upsertPartialRating(
            UpdatePartialRatingDto updatePartialRatingDto, FirebaseAgent.UserAuthData authentication) {
        log.info("Upserting (partial) rating {}", updatePartialRatingDto.getRatingId() == null ?
                updatePartialRatingDto.getPartialRatingId() : updatePartialRatingDto.getRatingId());
        var updatedRating = partialRatingRepository.upsertPartialRating(updatePartialRatingDto, authentication);
        return DtoMapper.INSTANCE.partialRatingToDtos(updatedRating);
    }
}
