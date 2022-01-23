from typing import Any, Dict, List, Type, TypeVar, Union, cast

import attr

from ..types import UNSET, Unset

T = TypeVar("T", bound="Template")


@attr.s(auto_attribs=True)
class Template:
    """
    Attributes:
        id (Union[Unset, int]):
        title (Union[Unset, str]):
        content (Union[Unset, str]):
        subject (Union[Unset, str]):
        is_being_edited_by (Union[Unset, int]):
        attachement_ids (Union[Unset, List[int]]):
    """

    id: Union[Unset, int] = UNSET
    title: Union[Unset, str] = UNSET
    content: Union[Unset, str] = UNSET
    subject: Union[Unset, str] = UNSET
    is_being_edited_by: Union[Unset, int] = UNSET
    attachement_ids: Union[Unset, List[int]] = UNSET
    additional_properties: Dict[str, Any] = attr.ib(init=False, factory=dict)

    def to_dict(self) -> Dict[str, Any]:
        id = self.id
        title = self.title
        content = self.content
        subject = self.subject
        is_being_edited_by = self.is_being_edited_by
        attachement_ids: Union[Unset, List[int]] = UNSET
        if not isinstance(self.attachement_ids, Unset):
            attachement_ids = self.attachement_ids

        field_dict: Dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if id is not UNSET:
            field_dict["id"] = id
        if title is not UNSET:
            field_dict["title"] = title
        if content is not UNSET:
            field_dict["content"] = content
        if subject is not UNSET:
            field_dict["subject"] = subject
        if is_being_edited_by is not UNSET:
            field_dict["isBeingEditedBy"] = is_being_edited_by
        if attachement_ids is not UNSET:
            field_dict["attachementIds"] = attachement_ids

        return field_dict

    @classmethod
    def from_dict(cls: Type[T], src_dict: Dict[str, Any]) -> T:
        d = src_dict.copy()
        id = d.pop("id", UNSET)

        title = d.pop("title", UNSET)

        content = d.pop("content", UNSET)

        subject = d.pop("subject", UNSET)

        is_being_edited_by = d.pop("isBeingEditedBy", UNSET)

        attachement_ids = cast(List[int], d.pop("attachementIds", UNSET))

        template = cls(
            id=id,
            title=title,
            content=content,
            subject=subject,
            is_being_edited_by=is_being_edited_by,
            attachement_ids=attachement_ids,
        )

        template.additional_properties = d
        return template

    @property
    def additional_keys(self) -> List[str]:
        return list(self.additional_properties.keys())

    def __getitem__(self, key: str) -> Any:
        return self.additional_properties[key]

    def __setitem__(self, key: str, value: Any) -> None:
        self.additional_properties[key] = value

    def __delitem__(self, key: str) -> None:
        del self.additional_properties[key]

    def __contains__(self, key: str) -> bool:
        return key in self.additional_properties
