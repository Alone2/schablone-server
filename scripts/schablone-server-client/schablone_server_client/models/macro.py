from typing import Any, Dict, List, Type, TypeVar, Union

import attr

from ..types import UNSET, Unset

T = TypeVar("T", bound="Macro")


@attr.s(auto_attribs=True)
class Macro:
    """
    Attributes:
        id (Union[Unset, int]):
        title (Union[Unset, str]):
        is_being_edited_by (Union[Unset, int]):
        content (Union[Unset, str]):
    """

    id: Union[Unset, int] = UNSET
    title: Union[Unset, str] = UNSET
    is_being_edited_by: Union[Unset, int] = UNSET
    content: Union[Unset, str] = UNSET
    additional_properties: Dict[str, Any] = attr.ib(init=False, factory=dict)

    def to_dict(self) -> Dict[str, Any]:
        id = self.id
        title = self.title
        is_being_edited_by = self.is_being_edited_by
        content = self.content

        field_dict: Dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if id is not UNSET:
            field_dict["id"] = id
        if title is not UNSET:
            field_dict["title"] = title
        if is_being_edited_by is not UNSET:
            field_dict["isBeingEditedBy"] = is_being_edited_by
        if content is not UNSET:
            field_dict["content"] = content

        return field_dict

    @classmethod
    def from_dict(cls: Type[T], src_dict: Dict[str, Any]) -> T:
        d = src_dict.copy()
        id = d.pop("id", UNSET)

        title = d.pop("title", UNSET)

        is_being_edited_by = d.pop("isBeingEditedBy", UNSET)

        content = d.pop("content", UNSET)

        macro = cls(
            id=id,
            title=title,
            is_being_edited_by=is_being_edited_by,
            content=content,
        )

        macro.additional_properties = d
        return macro

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
