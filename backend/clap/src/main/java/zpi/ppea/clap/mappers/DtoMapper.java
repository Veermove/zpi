package zpi.ppea.clap.mappers;

import data_store.*;
import org.mapstruct.Context;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.factory.Mappers;
import zpi.ppea.clap.dtos.*;

import java.util.List;
import java.util.Optional;

@Mapper
public interface DtoMapper {

    DtoMapper INSTANCE = Mappers.getMapper(DtoMapper.class);

    @Mapping(target = "partialRatings", source = "partialRatingsList")
    AssessorRatingsDto assessorRatingsToDtos(AssessorRatings assessorRatings);

    List<AssessorRatingsDto> assessorRatingsListToDtos(List<AssessorRatings> assessorRatings);

    CriterionDto criterionToDto(Criterion rating);

    List<CriterionDto> criterionListToDto(List<Criterion> rating);

    PartialRatingDto partialRatingToDtos(PartialRating partialRatings);

    List<PartialRatingDto> partialRatingListToDtos(List<PartialRating> partialRatings);

    @Mapping(target = "appReportDto", source = "report")
    DetailsSubmissionResponseDto detailsSubmissionToDto(DetailsSubmissionResponse response);

    AppReportDto appReportToDto(AppReport report);

    List<SubmissionDto> submissionListToDtos(List<Submission> submissions);

    @Mapping(target = "assessors", source = "assessorsList")
    @Mapping(target = "ratings", source = "ratingsList")
    SubmissionDto submissionToDto(Submission submission);

    RatingDto ratingToDto(Rating rating);

    PartialRatingRequest dtoToPartialRatingRequest(UpdatePartialRatingDto dto, @Context Integer assessorId);

    default RatingsSubmissionResponseDto ratingsResponseToDto(RatingsSubmissionResponse ratings) {
        return RatingsSubmissionResponseDto.builder()
                .criteria(DtoMapper.INSTANCE.criterionListToDto(
                        Optional.of(ratings.getCriteriaList()).orElse(List.of())
                ))
                .individualRatings(DtoMapper.INSTANCE.assessorRatingsListToDtos(
                        Optional.of(ratings.getIndividualList()).orElse(List.of())
                ))
                .initialRating(DtoMapper.INSTANCE.assessorRatingsToDtos(ratings.getInitial()))
                .finalRating(DtoMapper.INSTANCE.assessorRatingsToDtos(ratings.getInitial()))
                .build();
    }

}
